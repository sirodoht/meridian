package internal

import (
	"fmt"

	"go.uber.org/zap"

	"nimona.io"
)

type API interface {
	Register(req *RegisterRequest) (*RegisterResponse, error)
	Login(req *AuthenticateRequest) (*AuthenticateResponse, error)
	GetProfile(req *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(req *UpdateProfileRequest) (*UpdateProfileResponse, error)
	CreateNote(req *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNotes(req *GetNotesRequest) (*GetNotesResponse, error)
	GetNote(req *GetNoteRequest) (*GetNoteResponse, error)
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
		User        *User            `json:"user"`
		Identity    *nimona.Identity `json:"identity"`
		Keygraph    *nimona.KeyGraph `json:"keygraph"`
		KeysCurrent *nimona.KeyPair  `json:"keysCurrent"`
		KeysNext    *nimona.KeyPair  `json:"keysNext"`
	}
)

func (api *api) Register(req *RegisterRequest) (*RegisterResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
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

func (api *api) Login(req *AuthenticateRequest) (*AuthenticateResponse, error) {
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

func (api *api) GetProfile(req *GetProfileRequest) (*GetProfileResponse, error) {
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

func (api *api) UpdateProfile(req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
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

func (api *api) CreateNote(req *CreateNoteRequest) (*CreateNoteResponse, error) {
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

func (api *api) GetNotes(req *GetNotesRequest) (*GetNotesResponse, error) {
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

func (api *api) GetNote(req *GetNoteRequest) (*GetNoteResponse, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}
