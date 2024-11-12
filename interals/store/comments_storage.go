package store

import (
	"context"
	"database/sql"
)

type ICommentsRepo interface {
	Get(context.Context, string) (*Users, error)
	Create(context.Context, Users) error
	Update(context.Context, string, Users) error
	Delete(context.Context, string) error
}

type CommentsRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentsRepo {
	return &CommentsRepo{db: db}
}
func (p *CommentsRepo) Get(context.Context, string) (*Users, error) {

	return nil, nil
}

func (p *CommentsRepo) Create(context.Context, Users) error {
	return nil
}
func (p *CommentsRepo) Update(context.Context, string, Users) error {
	return nil
}
func (p *CommentsRepo) Delete(context.Context, string) error {
	return nil
}
