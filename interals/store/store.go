package store

import (
	"database/sql"
)

type Storage interface {
	Posts() IPostsRepo
	Comments() ICommentsRepo
	Users() IUsersRepo
}

type PostgresStore struct {
	db       *sql.DB
	Users    IUsersRepo
	Posts    IPostsRepo
	Comments ICommentsRepo
}

func (s *PostgresStore) NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db, Posts: NewPostRepo(db),
		Users:    NewUsersRepo(db),
		Comments: NewCommentRepo(db)}
}
