package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app/models"
	"github.com/jasurxaydarov/todo_app/storage/repoi"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(con *pgx.Conn)repoi.UserRepoI{

	return &UserRepo{
		conn: con,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error                 {

	return nil
}
func (u *UserRepo) GetUSer(ctx context.Context, limit int, page int) (*models.User, error) {

	return nil,nil
}
