package routers

import (

	"github.com/gorilla/mux"
	"github.com/wjurmy/havaard/controllers"
)

/*
			Routes for the Users Resource in user.go
			----------------------------------------
	SetUserRoutes function receives a pointer to the Gorilla mux
	router object ( mux.Router ) as an argument pointer of the
	mux.Router object. Two routes are specified: for registering
	a new user login to the system. Application handler functions
	called from the controllers package.*/

func SetUserRoutes(router *mux.Router) *mux.Router{
	router.HandleFunc("/users/registration", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login",controllers.Login).Methods("POST")
	router.HandleFunc("/users/{id}",controllers.GetUserByID).Methods("GET")


	return router
}
