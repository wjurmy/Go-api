package main

import (
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"github.com/wjurmy/havaard/common"
	"github.com/wjurmy/havaard/routers"

	//"github.com/gorilla/handlers"
)

func init(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Main function...")

	// Calls start up logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	/*router.Walk( func (route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		log.Println(path)
		return nil
	})*/



	// Create negroni instance
	n := negroni.Classic()


	handlers := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:27017"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT","GET","POST", "DELETE","*","PATCH"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "*","token"},
		Debug:            true,
	}).Handler(router)

	n.UseHandler(handlers)

	server := &http.Server{
		Addr:    ":"+common.AppConfig.Server,
		Handler: n,
	}


	log.Print("Listening on port..", common.AppConfig.Server)
	log.Fatal(server.ListenAndServe())
}
