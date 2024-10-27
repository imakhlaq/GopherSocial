package store

import (
	"context"
	"database/sql"
)

type Users struct {
	Id       string `jason:"id"`
	Username string `jason:"username"`
	Password string `jason:"-"`
	Email    string `jason:"email"`
}

// anystoreage need to immplement these methods
type IUsersRepo interface {
	Get(context.Context, string) (*Users, error)
	Create(context.Context, Users) error
	Update(context.Context, string, Users) error
	Delete(context.Context, string) error
}

// implementing above methods
type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) IUsersRepo {
	return &UsersRepo{db: db}
}

func (u *UsersRepo) Get(ctx context.Context, uuid string) (*Users, error) {

}

func (u *UsersRepo) Create(ctx context.Context, user Users) error {

}

func (u *UsersRepo) Update(ctx context.Context, uuid string, update Users) error {

}

func (u *UsersRepo) Delete(ctx context.Context, uuid string) error {

}
