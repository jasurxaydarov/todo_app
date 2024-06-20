package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jasurxaydarov/todo_app/api"
	"github.com/jasurxaydarov/todo_app/api/handlers"
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

		log.Println("error on main storege ",storage)
		return
	}

	log.Println("sucessfully got storage ")

	fmt.Println(storage)

	options:=api.Options{

		Cfg:cfg,
		HandlerOPtions: handlers.Handlers{

			Storage:storage,

		},
	
	
		}
	
	api:=api.Api(options)

	server:=http.Server{

		Addr: cfg.HttpPort,
		Handler: api,

	}
	
	log.Println("server ready for serve")
	log.Println("server running on :",cfg.HttpPort)

	if err:=server.ListenAndServe();err!=nil{

		log.Println("error on listen and serve")
		return
	}

}
