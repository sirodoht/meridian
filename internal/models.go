package internal

import (
	"time"

	"nimona.io"
)

type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Document struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Session struct {
	ID        int64  `db:"id"`
	UserID    int64  `db:"user_id"`
	TokenHash string `db:"token_hash"`
}

type (
	NimonaFeed struct {
		Metadata nimona.Metadata `nimona:"$metadata,type=feed"`
		Profile  NimonaProfile   `nimona:"profile,omitempty"`
		Notes    []*NimonaNote   `nimona:"posts,omitempty"`
		Folowees []*NimonaFollow `nimona:"folowees,omitempty"`
	}
	NimonaProfile struct {
		Metadata    nimona.Metadata `nimona:"$metadata,type=profile"`
		DisplayName string          `nimona:"displayName,omitempty"`
		Description string          `nimona:"description,omitempty"`
		AvatarURL   string          `nimona:"avatarURL,omitempty"`
	}
	NimonaNote struct {
		Metadata  nimona.Metadata `nimona:"$metadata,type=post"`
		Key       string          `nimona:"_key,omitempty"`
		Partition string          `nimona:"_partition,omitempty"`
		Content   string          `nimona:"content"`
	}
	NimonaFollow struct {
		Metadata  nimona.Metadata      `nimona:"$metadata,type=follow"`
		Identity  nimona.Identity      `nimona:"identity,omitempty"`
		Alias     nimona.IdentityAlias `nimona:"alias,omitempty"`
		Timestamp string               `nimona:"timestamp,omitempty"`
	}
)
