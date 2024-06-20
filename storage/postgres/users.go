package postgres

import (
	"context"
	"log"

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

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error{

	query := `
		INSERT INTO 
			users (
				user_id, 
				user_name, 
				password, 
				created_at
			) VALUES (
			 	$1, $2, $3, $4
			)`
	_,err:=u.conn.Exec(
		ctx,query,
		user.UserID,
		user.UserName,
		user.Password,
		user.CreatedAt,
	
	)

	if err != nil {


		log.Println("error on CreateUser",err)
		return err
	}
	

	return nil
}
func (u *UserRepo) GetUSer(ctx context.Context, limit int, page int) (*models.User, error) {

	return nil,nil
}
