package repoi

import (
	"context"

	"github.com/jasurxaydarov/todo_app/models"
)

type UserRepoI interface{
	
	CreateUser(ctx context.Context, user models.User)error
	GetUSer(ctx context.Context, limit, page int)(*models.User, error)
}