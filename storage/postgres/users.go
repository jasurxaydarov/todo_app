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
func (u *UserRepo) GetUSers(ctx context.Context, limit int, page int) ([]*models.User, error) {
	
	var users []*models.User
	var user models.User

	query:=`
		SELECT
   			 *
		FROM
    		users

		LIMIT $1

	 	OFFSET $2


	`

	row, err := u.conn.Query(ctx,query,limit,page)

	if err != nil{

		log.Println("err on GetUSerssssssss",err)
		return nil,nil
	}

	for row.Next(){

		err:=row.Scan(
			&user.UserID,
			&user.UserName,
			&user.Password,
			&user.CreatedAt,
		)

		if err != nil{

			log.Println("err on GetUSers scan",err)
			return nil,nil
		}

		users=append(users, &user)
		
	}

	defer row.Close()
	return users,nil

}

func(u *UserRepo)GetUSerById(ctx context.Context,name string)(*models.User,error){

	var user models.User

	query:=`
		SELECT 
			*
		FROM
			users
		WHERE
			user_name=$1 
	`

	err:=u.conn.QueryRow(
		ctx,query,
		name,
		).Scan(
			&user.UserID,
			&user.UserName,
			&user.Password,
			&user.CreatedAt,
		)

		if err!= nil{

			log.Println("err on aaaaaaa GetUSerById",err)
			return nil,err

		}
		


	return &user,nil

} 
