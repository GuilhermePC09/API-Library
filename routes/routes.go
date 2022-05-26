package routes

import (
	"github.com/GuilhermePC09/API-Library/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()

	// Users routes
	router.HandleFunc("/users/registration", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/status/{userId}", controllers.ChangeUserStatus).Methods("PUT")
	router.HandleFunc("/users/edit", controllers.EditUser).Methods("PUT")
	router.HandleFunc("/users/delete/{userId}", controllers.EditUser).Methods("DELETE")

	// Books routes
	router.HandleFunc("/books/registration", controllers.RecordBook).Methods("POST")
	router.HandleFunc("/books", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/books/rental/{bookId}", controllers.RentBook).Methods("GET")
	router.HandleFunc("/books/return/{bookId}", controllers.ReturnBook).Methods("GET")
	router.HandleFunc("/books/edit", controllers.EditBook).Methods("PUT")
	router.HandleFunc("/users/delete/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
