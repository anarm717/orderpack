package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Println("\nService started. Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
