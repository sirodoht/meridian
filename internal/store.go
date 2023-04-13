package internal

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"nimona.io"
)

var ErrNotFound = fmt.Errorf("not found")

type Store interface {
	PutUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, error)
	PutProfile(context.Context, *Profile) error
	GetProfile(context.Context, string) (*Profile, error)
	PutNote(context.Context, *Note) error
	GetNotes(context.Context, string, int, int) ([]*Note, error)
	PutFollow(context.Context, *Follow) error
	GetFollowers(context.Context, nimona.KeygraphID) ([]*Follow, error)
	GetFollowees(context.Context, nimona.KeygraphID) ([]*Follow, error)
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
		&Follow{},
		&Note{},
		&Session{},
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
	username string,
) (*User, error) {
	if username == "" {
		return nil, fmt.Errorf("failed to get user: nil request")
	}

	var user User
	err := s.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// TODO: improve not found check
	if user.KeygraphID.IsEmpty() {
		return nil, ErrNotFound
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
		return nil, ErrNotFound
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
	offset int,
	limit int,
) ([]*Note, error) {
	query := s.db.
		WithContext(ctx).
		Preload("Profile")

	if identityNRI != "" {
		query = query.Where("identity = ?", identityNRI)
	}

	var notes []*Note
	err := query.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&notes).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}

	return notes, nil
}

func (s *SQLStore) PutFollow(
	ctx context.Context,
	req *Follow,
) error {
	if req == nil {
		return fmt.Errorf("failed to put follow: nil request")
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
		return fmt.Errorf("failed to put follow: %w", err)
	}

	return nil
}

func (s *SQLStore) GetFollowers(
	ctx context.Context,
	followee nimona.KeygraphID,
) ([]*Follow, error) {
	if followee.IsEmpty() {
		return nil, fmt.Errorf("failed to get followers: nil request")
	}

	var follows []*Follow
	err := s.db.
		WithContext(ctx).
		Where("followee = ?", followee).
		Find(&follows).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get follows: %w", err)
	}

	return follows, nil
}

func (s *SQLStore) GetFollowees(
	ctx context.Context,
	follower nimona.KeygraphID,
) ([]*Follow, error) {
	if follower.IsEmpty() {
		return nil, fmt.Errorf("failed to get followees: nil request")
	}

	var follows []*Follow
	err := s.db.
		WithContext(ctx).
		Where("follower = ?", follower).
		Find(&follows).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get follows: %w", err)
	}

	return follows, nil
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
	id string,
) (*Session, error) {
	if id == "" {
		return nil, fmt.Errorf("nil request")
	}

	var session Session
	err := s.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&session).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	// TODO: improve not found check
	if session.ID == "" {
		return nil, ErrNotFound
	}

	return &session, nil
}
