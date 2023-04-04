package internal

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store interface {
	PutUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, error)
	PutProfile(context.Context, *Profile) error
	GetProfile(context.Context, string) (*Profile, error)
	PutNote(context.Context, *Note) error
	GetNotes(context.Context, string) ([]*Note, error)
	PutSession(context.Context, *Session) error
	GetSession(context.Context, string) (*Session, error)
}

type SQLStore struct {
	db *gorm.DB
}

func NewSQLStore(gdb *gorm.DB) Store {
	gdb.AutoMigrate(
		&User{},
		&Profile{},
		&Note{},
	)

	return &SQLStore{
		db: gdb,
	}
}

// PutUser will create or update a user
func (s *SQLStore) PutUser(
	ctx context.Context,
	req *User,
) error {
	if req == nil {
		return fmt.Errorf("failed to put user: nil request")
	}

	err := s.db.
		WithContext(ctx).
		Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).
		Create(req).
		Error
	if err != nil {
		return fmt.Errorf("failed to put user: %w", err)
	}

	return nil
}

func (s *SQLStore) GetUser(
	ctx context.Context,
	identityNRI string,
) (*User, error) {
	if identityNRI == "" {
		return nil, fmt.Errorf("failed to get user: nil request")
	}

	var user User
	err := s.db.
		WithContext(ctx).
		Where("id = ?", identityNRI).
		First(&user).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// TODO: improve not found check
	if user.IdentityNRI == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (s *SQLStore) PutProfile(
	ctx context.Context,
	req *Profile,
) error {
	if req == nil {
		return fmt.Errorf("failed to put profile: nil request")
	}

	err := s.db.
		WithContext(ctx).
		Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).
		Create(req).
		Error
	if err != nil {
		return fmt.Errorf("failed to put profile: %w", err)
	}

	return nil
}

func (s *SQLStore) GetProfile(
	ctx context.Context,
	identityNRI string,
) (*Profile, error) {
	if identityNRI == "" {
		return nil, fmt.Errorf("failed to get profile: nil request")
	}

	var profile Profile
	err := s.db.
		WithContext(ctx).
		Where("id = ?", identityNRI).
		First(&profile).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &profile, nil
}

func (s *SQLStore) PutNote(
	ctx context.Context,
	req *Note,
) error {
	if req == nil {
		return fmt.Errorf("failed to put note: nil request")
	}

	err := s.db.
		WithContext(ctx).
		Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).
		Create(req).
		Error
	if err != nil {
		return fmt.Errorf("failed to put note: %w", err)
	}

	return nil
}

func (s *SQLStore) GetNotes(
	ctx context.Context,
	identityNRI string,
	// TODO: add filters
) ([]*Note, error) {
	if identityNRI == "" {
		return nil, fmt.Errorf("failed to get notes: nil request")
	}

	var notes []*Note
	err := s.db.
		WithContext(ctx).
		Where("id = ?", identityNRI).
		Find(&notes).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}

	return notes, nil
}

func (s *SQLStore) PutSession(
	ctx context.Context,
	req *Session,
) error {
	if req == nil {
		return fmt.Errorf("failed to put session: nil request")
	}

	err := s.db.
		WithContext(ctx).
		Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).
		Create(req).
		Error
	if err != nil {
		return fmt.Errorf("failed to put session: %w", err)
	}

	return nil
}

func (s *SQLStore) GetSession(
	ctx context.Context,
	token string,
) (*Session, error) {
	if token == "" {
		return nil, fmt.Errorf("nil request")
	}

	var session Session
	err := s.db.
		WithContext(ctx).
		Where("token = ?", token).
		First(&session).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	// TODO: improve not found check
	if session.Token == "" {
		return nil, fmt.Errorf("session not found")
	}

	return &session, nil
}
