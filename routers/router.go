package routers

import (
	"github.com/gorilla/mux"

)
// The InitRoutes function in main.go is called when the
// HTTP server starts

func InitRoutes() *mux.Router {
	router :=mux.NewRouter().StrictSlash(false)

	// Routes for the user entity
	router = SetUserRoutes(router)

	// Routes for the company entity
	router = SetCompanyRoutes(router)


	return router
}