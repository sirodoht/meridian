package internal

import (
	"time"

	"nimona.io"
)

type (
	User struct {
		IdentityNRI  string `gorm:"type:varchar(255);primary_key"`
		Username     string `gorm:"type:varchar(255);unique,index"`
		Email        string `gorm:"type:varchar(255);unique"`
		PasswordHash string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
	Profile struct {
		IdentityNRI string `gorm:"type:varchar(255);primary_key"`
		DisplayName string
		Description string
		AvatarURL   string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	Note struct {
		IdentityNRI string
		NoteID      string `gorm:"type:varchar(255);primary_key"`
		Content     string `gorm:"type:text"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

type (
	Session struct {
		ID        string `gorm:"type:varchar(255);primary_key"`
		Username  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

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
