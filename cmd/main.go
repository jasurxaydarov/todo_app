package main

import (
	"fmt"
	"log"

	"github.com/jasurxaydarov/todo_app/config"
	"github.com/jasurxaydarov/todo_app/pgx"
	"github.com/jasurxaydarov/todo_app/storage"
)

func main() {

	cfg := config.Load()

	conn, err:=pgx.NewPgx(*cfg)

	if err != nil {

		return		
	}

	log.Println("sucessfully connected with db ")

	fmt.Println(conn)

	storage:=storage.NewStorage(conn)

	if storage == nil{
		return
	}

	log.Println("sucessfully got storage ")

	fmt.Println(storage)

}
