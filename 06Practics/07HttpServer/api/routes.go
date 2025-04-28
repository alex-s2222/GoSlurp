package api

import (
	"httpServer/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UserHandler) *mux.Router{

	r := mux.NewRouter()
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")
}