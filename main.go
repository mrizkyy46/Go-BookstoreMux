package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrizkyy46/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
