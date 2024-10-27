package store

import (
	"context"
	"database/sql"
)

type IPostsRepo interface {
	Get(context.Context, string) (*Users, error)
	Create(context.Context, Users) error
	Update(context.Context, string, Users) error
	Delete(context.Context, string) error
}
type Posts struct {
	ID      string `json:"id"`
	Tittle  string `json:"tittle"`
	Content string `json:"content"`
}

type PostsRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) IPostsRepo {
	return &PostsRepo{db: db}
}
func (p *PostsRepo) Get(context.Context, string) (*Users, error) {
	query := "SELECT * FROM users WHERE id = %s"

}

func (p *PostsRepo) Create(context.Context, Users) error {

}
func (p *PostsRepo) Update(context.Context, string, Users) error {

}
func (p *PostsRepo) Delete(context.Context, string) error {

}
