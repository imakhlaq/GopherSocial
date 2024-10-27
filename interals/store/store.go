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
	db *sql.DB
}

// Comments implements Storage.
func (p PostgresStore) Comments() ICommentsRepo {
	return NewCommentRepo(p.db)
}

// Posts implements Storage.
func (p PostgresStore) Posts() IPostsRepo {
	return NewPostRepo(p.db)
}

// Users implements Storage.
func (p PostgresStore) Users() IUsersRepo {
	return NewUsersRepo(p.db)
}

func NewStore(db *sql.DB) Storage {
	return PostgresStore{db: db}
}
