package internal

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"nimona.io"
)

type API interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error)
	GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error)
}

func NewAPI(
	logger *zap.Logger,
	meridianStore Store,
	documentStore *nimona.DocumentStore,
	identityStore *nimona.IdentityStore,
) API {
	return &api{
		logger:        logger,
		meridianStore: meridianStore,
		documentStore: documentStore,
		identityStore: identityStore,
	}
}

type api struct {
	logger        *zap.Logger
	meridianStore Store
	documentStore *nimona.DocumentStore
	identityStore *nimona.IdentityStore
}

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	RegisterResponse struct {
		User      *User            `json:"user"`
		Identity  *nimona.Identity `json:"identity"`
		SessionID string           `json:"sessionId"`
	}
)

func (api *api) Register(
	ctx context.Context,
	req *RegisterRequest,
) (*RegisterResponse, error) {
	// create new identity
	// TODO: add missing use, once NRI support for use is added
	id, err := api.identityStore.NewIdentity("")
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
	}

	// store identity
	err = api.documentStore.PutDocument(id.Document())
	if err != nil {
		return nil, fmt.Errorf("failed to put identity: %w", err)
	}

	// create new feed
	feed := &NimonaFeed{
		Metadata: nimona.Metadata{
			Owner: id,
		},
	}

	// store note feed
	err = api.documentStore.PutDocument(feed.Document())
	if err != nil {
		return nil, fmt.Errorf("failed to put note feed: %w", err)
	}

	// hash password
	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// create user
	user := &User{
		IdentityNRI:  id.String(),
		Username:     req.Username,
		PasswordHash: passwordHash,
		Email:        req.Email,
	}
	err = api.meridianStore.PutUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to put user: %w", err)
	}

	// update profile
	_, err = api.UpdateProfile(ctx, &UpdateProfileRequest{
		IdentityNRI: id.String(),
		DisplayName: req.Username,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	// create session
	ses := &Session{
		ID:       uuid.NewString(),
		Username: req.Username,
	}
	err = api.meridianStore.PutSession(ctx, ses)
	if err != nil {
		return nil, fmt.Errorf("failed to put session: %w", err)
	}

	res := &RegisterResponse{
		User:      user,
		Identity:  id,
		SessionID: ses.ID,
	}
	return res, nil
}

type (
	AuthenticateRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	AuthenticateResponse struct {
		SessionID string `json:"sessionId"`
		User      *User  `json:"user"`
	}
)

func (api *api) Login(
	ctx context.Context,
	req *AuthenticateRequest,
) (*AuthenticateResponse, error) {
	// get user
	user, err := api.meridianStore.GetUser(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// create session
	ses := &Session{
		ID:       uuid.NewString(),
		Username: req.Username,
	}
	err = api.meridianStore.PutSession(ctx, ses)
	if err != nil {
		return nil, fmt.Errorf("failed to put session: %w", err)
	}

	return &AuthenticateResponse{
		SessionID: ses.ID,
		User:      user,
	}, nil
}

type (
	GetProfileRequest struct {
		IdentityNRI string `json:"identity"`
	}
	GetProfileResponse struct {
		Identity *nimona.Identity `json:"identity"`
		Profile  *NimonaProfile   `json:"profile"`
	}
)

func (api *api) GetProfile(
	ctx context.Context,
	req *GetProfileRequest,
) (*GetProfileResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}

type (
	UpdateProfileRequest struct {
		IdentityNRI string `json:"identity"`
		DisplayName string `json:"displayName,omitempty"`
		Description string `json:"description,omitempty"`
		AvatarURL   string `json:"avatarUrl,omitempty"`
	}
	UpdateProfileResponse struct{}
)

func (api *api) UpdateProfile(
	ctx context.Context,
	req *UpdateProfileRequest,
) (*UpdateProfileResponse, error) {
	// get identity
	id, err := nimona.ParseIdentityNRI(req.IdentityNRI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse identity: %w", err)
	}

	// figure out feed root id
	feed := &NimonaFeed{
		Metadata: nimona.Metadata{
			Owner: id,
		},
	}
	feedRootID := nimona.NewDocumentID(feed.Document())

	// get signing context
	sctx, err := api.getSigningContext(req.IdentityNRI)
	if err != nil {
		return nil, fmt.Errorf("failed to get signing context: %w", err)
	}

	// create profile
	profile := &NimonaProfile{
		Metadata: nimona.Metadata{
			Owner: id,
		},
		DisplayName: req.DisplayName,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
	}

	patchDoc, err := api.documentStore.CreatePatch(
		feedRootID,
		"replace",
		"profile",
		profile.Map(),
		*sctx,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create patch: %w", err)
	}

	// store patch
	err = api.documentStore.PutDocument(patchDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to put profile: %w", err)
	}

	return &UpdateProfileResponse{}, nil
}

type (
	CreateNoteRequest struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}
	CreateNoteResponse struct{}
)

func (api *api) CreateNote(
	ctx context.Context,
	req *CreateNoteRequest,
) (*CreateNoteResponse, error) {
	// get user
	user, err := api.meridianStore.GetUser(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// get identity
	id, err := nimona.ParseIdentityNRI(user.IdentityNRI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse identity: %w", err)
	}
	kg, err := api.identityStore.IdentityStore.Get(*id)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity: %w", err)
	}
	kp, err := api.identityStore.KeyPairStore.Get(kg.Keys)
	if err != nil {
		return nil, fmt.Errorf("failed to get keypair: %w", err)
	}

	// get feed for identity
	// TODO: get feed from identity

	// create note
	// TODO: add metadata
	note := &NimonaNote{
		Content: req.Content,
	}
	// TODO: sign note
	// TODO: create patch
	// TODO: store document
	// TODO: apply patch

	fmt.Println(kp, note)
	return nil, fmt.Errorf("not implemented")
}

type (
	GetNotesRequest struct {
		IdentityNRI string `json:"identity"`
		// TODO: add filters
		// TODO: add pagination
	}
	GetNotesResponse struct {
		Notes []*Note `json:"notes"`
	}
)

func (api *api) GetNotes(
	ctx context.Context,
	req *GetNotesRequest,
) (*GetNotesResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}

type (
	GetNoteRequest struct {
		NoteID string `json:"noteId"`
	}
	GetNoteResponse struct {
		Note *Note `json:"note"`
	}
)

func (api *api) GetNote(
	ctx context.Context,
	req *GetNoteRequest,
) (*GetNoteResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
