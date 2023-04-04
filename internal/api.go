package internal

import (
	"context"
	"fmt"

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
	}
	RegisterResponse struct {
		User     *User            `json:"user"`
		Identity *nimona.Identity `json:"identity"`
	}
)

func (api *api) Register(
	ctx context.Context,
	req *RegisterRequest,
) (*RegisterResponse, error) {
	// create new identity
	id, err := api.identityStore.NewIdentity("user")
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
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
	}
	err = api.meridianStore.PutUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to put user: %w", err)
	}

	res := &RegisterResponse{
		User:     user,
		Identity: id,
	}
	return res, nil
}

type (
	AuthenticateRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	AuthenticateResponse struct {
		Token string `json:"token"`
	}
)

func (api *api) Login(
	ctx context.Context,
	req *AuthenticateRequest,
) (*AuthenticateResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
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
	UpdateProfileResponse struct {
		Identity *nimona.Identity `json:"identity"`
		Profile  *NimonaProfile   `json:"profile"`
	}
)

func (api *api) UpdateProfile(
	ctx context.Context,
	req *UpdateProfileRequest,
) (*UpdateProfileResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}

type (
	CreateNoteRequest struct {
		IdentityNRI string `json:"identity"`
		Content     string `json:"content"`
	}
	CreateNoteResponse struct{}
)

func (api *api) CreateNote(
	ctx context.Context,
	req *CreateNoteRequest,
) (*CreateNoteResponse, error) {
	// TODO: implement
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
