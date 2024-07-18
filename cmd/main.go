package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/gorilla/handlers"
	"github.com/4lerman/proxy_server/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/proxy", handlers.HandlerProxy).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))

}
