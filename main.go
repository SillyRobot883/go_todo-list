package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// init DB
	InitDB()

	// init router
	r := mux.NewRouter()

	prefix := r.PathPrefix("/api/v1").Subrouter()

	// @TODO - handler routes
	// get all the tasks
	prefix.HandleFunc("/task", getTasks).Methods("GET")
	// get one task
	prefix.HandleFunc("/task/{id}", getTask).Methods("GET")
	// create a task
	prefix.HandleFunc("/task", createTask).Methods("POST")
	// update a task
	prefix.HandleFunc("/task/{id}", updateTask).Methods("PATCH")
	// delete a task
	prefix.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	// Add CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Adjust this to only allow specific origins
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Fatal(http.ListenAndServe(":8085", corsHandler(r)))
}
