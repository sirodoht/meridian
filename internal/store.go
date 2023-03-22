package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type SQLStore struct {
	db *sqlx.DB
}

func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) InsertUser(ctx context.Context, d *User) (int64, error) {
	var id int64
	rows, err := s.db.NamedQuery(`
		INSERT INTO users (
			username,
			email,
			created_at,
			updated_at
		) VALUES (
			:username,
			:email,
			:created_at,
			:updated_at
		) RETURNING id`, d)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
	}
	return id, nil
}

func (s *SQLStore) InsertUserPage(
	ctx context.Context,
	username string,
	email string,
	passwordHash string,
) (int64, error) {
	var id int64
	timenow := time.Now()
	row := s.db.QueryRow(`
		INSERT INTO users (
			email,
			username,
			password_hash,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id`,
		email,
		username,
		passwordHash,
		timenow,
		timenow,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SQLStore) UpdateUser(
	ctx context.Context,
	id int64,
	field string,
	value string,
) error {
	sql := fmt.Sprintf("UPDATE users SET %s=:value WHERE id=:id", field)
	_, err := s.db.NamedExec(sql, map[string]interface{}{
		"field": field,
		"value": value,
		"id":    id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) GetOneUser(ctx context.Context, id int64) (*User, error) {
	var users []*User
	err := s.db.SelectContext(
		ctx,
		&users,
		`SELECT * FROM users WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func (s *SQLStore) GetOneUserByUsername(
	ctx context.Context,
	username string,
) (*User, error) {
	var users []*User
	err := s.db.SelectContext(
		ctx,
		&users,
		`SELECT * FROM users WHERE username=$1`,
		username,
	)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no user exists with this username")
	}
	return users[0], nil
}

func (s *SQLStore) InsertDocument(
	ctx context.Context,
	d *Document,
) (int64, error) {
	var id int64
	rows, err := s.db.NamedQuery(`
		INSERT INTO documents (
			title,
			body,
			created_at,
			updated_at
		) VALUES (
			:title,
			:body,
			:created_at,
			:updated_at
		) RETURNING id`, d)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
	}
	return id, nil
}

func (s *SQLStore) UpdateDocument(
	ctx context.Context,
	id int64,
	field string,
	value string,
) error {
	sql := fmt.Sprintf(`
		UPDATE documents
		SET
			%s=:value,
			updated_at=now()
		WHERE id=:id
	`, field)
	_, err := s.db.NamedExec(sql, map[string]interface{}{
		"field": field,
		"value": value,
		"id":    id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) GetAllDocument(ctx context.Context) ([]*Document, error) {
	var docs []*Document
	err := s.db.SelectContext(
		ctx,
		&docs,
		`SELECT * FROM documents ORDER BY title ASC`,
	)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func (s *SQLStore) GetOneDocument(
	ctx context.Context,
	id int64,
) (*Document, error) {
	var docs []*Document
	err := s.db.SelectContext(
		ctx,
		&docs,
		`SELECT * FROM documents WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return docs[0], nil
}

func (s *SQLStore) InsertSession(
	ctx context.Context,
	d *Session,
) (int64, error) {
	var id int64
	rows, err := s.db.NamedQuery(`
		INSERT INTO sessions (
			user_id,
			token_hash
		) VALUES (
			:user_id,
			:token_hash
		) RETURNING id`, d)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
	}
	return id, nil
}

func (s *SQLStore) GetOneSession(ctx context.Context, tokenHash string) (
	*Session,
	error,
) {
	var sessions []*Session
	err := s.db.SelectContext(
		ctx,
		&sessions,
		`SELECT * FROM sessions WHERE token_hash=$1`,
		tokenHash,
	)
	if err != nil {
		return nil, err
	}
	if len(sessions) == 0 {
		return nil, fmt.Errorf("no user exists with this username")
	}
	return sessions[0], nil
}

func (s *SQLStore) GetUsernameSession(
	ctx context.Context,
	tokenHash string,
) string {
	type UserSession struct {
		Username string
	}
	var userSessions []*UserSession
	err := s.db.SelectContext(
		ctx,
		&userSessions,
		`SELECT users.username
		FROM sessions JOIN users ON sessions.user_id = users.id
		WHERE token_hash=$1`,
		tokenHash,
	)
	if err != nil {
		panic(err)
	}
	if len(userSessions) == 0 {
		return ""
	}
	return userSessions[0].Username
}

func (s *SQLStore) DeleteSession(ctx context.Context, tokenHash string) error {
	_, err := s.db.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1`,
		tokenHash,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
