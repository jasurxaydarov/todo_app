package pgx

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/todo_app/config"
)

func NewPgx( cfg config.Config)(*pgx.Conn,error){

	url:=fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DbUser,
		cfg.DbUserPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
		

	)

	conn,err:=pgx.Connect(context.Background(),url)

	if err!=nil{

		log.Println("err on conn with db",err)
		return nil,err
	}

	return conn,nil
}