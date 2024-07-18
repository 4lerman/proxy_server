package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/4lerman/proxy_server/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/4lerman/proxy_server/internal/handlers"
	"github.com/gorilla/mux"
)

// @title Proxy Server
// @version 1.0
// @description This is a proxy server api
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	subRouter := router.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/proxy", handlers.HandlerProxy).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))

}
