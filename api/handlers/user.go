package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jasurxaydarov/todo_app/models"
)

func (h *handlers) CreateUsers(resp http.ResponseWriter, req *http.Request) {
	
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
 
	if err != nil {

		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	 user.UserID=uuid.New()

	user.CreatedAt = time.Now()

	err = h.storage.GetUserRepo().CreateUser(context.Background(), user)

	if err != nil {

		http.Error(resp, "err creating user ", http.StatusInternalServerError)
		return

	}

	resp.Header().Set("Content-Type", "application/json")

	json.NewEncoder(resp).Encode(user)


}

func (h *handlers)GetUSerById(resp http.ResponseWriter,req *http.Request){


	user,err:=h.storage.GetUserRepo().GetUSerById(context.Background(),"jasur1")

	if err != nil {

		http.Error(resp, "err  GetUSerById ", http.StatusInternalServerError)
		return

	}

	resp.Header().Set("Content-Type", "application/json")

	json.NewEncoder(resp).Encode(user)

}


func (h *handlers)GetUSers(resp http.ResponseWriter, req *http.Request){

	limi,page:= 3, 1

	user, err:= h.storage.GetUserRepo().GetUSers(context.Background(), limi, page)

	if err!= nil{

		http.Error(resp, "error on get users", http.StatusInternalServerError)

		return 
	}
	
	resp.Header().Set("Content-Type","application/json")
	
	json.NewEncoder(resp).Encode(user)

}
