package main

import (
	"fmt"

	"github.com/jasurxaydarov/todo_app/config"
	"github.com/jasurxaydarov/todo_app/pgx"
)

func main() {

	cfg := config.Load()

	conn, err:=pgx.NewPgx(*cfg)

	if err != nil {

		return		
	}

	fmt.Println(conn)

}
