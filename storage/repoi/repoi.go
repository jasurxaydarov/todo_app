package repoi

import (
	"context"
	"github.com/jasurxaydarov/todo_app/models"
)

type UserRepoI interface{
	
	CreateUser(ctx context.Context, user models.User)error
	GetUSerById(ctx context.Context,name string)(*models.User, error)
	UpdateUserById(ctx context.Context,user models.User)error
	DeleteUserByName(ctx context.Context,name string)error
	GetUSers(ctx context.Context, limit, page int)([]*models.User, error)
}