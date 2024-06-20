package api

import (
	"net/http"

	"github.com/jasurxaydarov/todo_app/api/handlers"
	"github.com/jasurxaydarov/todo_app/config"
)

type Options struct {
	Cfg            *config.Config
	HandlerOPtions handlers.Handlers
}


func Api(options Options)*http.ServeMux{

	h:= handlers.NewHadlers(options.HandlerOPtions)
	
	api:=http.NewServeMux()

	api.HandleFunc("POST /api/crete-user",h.CreateUsers)
	api.HandleFunc("GET /api/get-user",h.GetUSerById)
	api.HandleFunc("GET /api/get-users",h.GetUSers)

	return api
}