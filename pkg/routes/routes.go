package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrizkyy46/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/create/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteBook).Methods("DELETE")
}
