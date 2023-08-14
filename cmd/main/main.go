package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vladsw764/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
