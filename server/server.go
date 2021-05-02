package mockserver

import (
	"fmt"
	"log"
	"net/http"

	localDatabase "mockserver/database"

	"github.com/gorilla/mux"
)

var database *localDatabase.Database

func welcomePage(w http.ResponseWriter, req *http.Request) {
	// TODO: return graphical UI.
	// TODO: include data from database.
	fmt.Fprintf(w, "This is under progress.\n")
}
func getResponseFromDB(w http.ResponseWriter, req *http.Request) {
	// TODO: parse db data and generate response.
	fmt.Fprintf(w, "This is under progress2.\n")
}

// StartServer is main entry point.
func StartServer(address string, db *localDatabase.Database) {
	database = db
	r := mux.NewRouter()
	r.HandleFunc("/", welcomePage).Methods("GET")
	r.HandleFunc("/get", get).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", getByID).Methods("GET")
	r.HandleFunc("/add", add).Methods("POST")
	r.HandleFunc("/update", update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", delete).Methods("DELETE")
	r.HandleFunc("/{path:[\\w\\W]+}", getResponseFromDB)

	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}
	log.Printf("Starting server at: '%s'", address)
	srv.ListenAndServe()
}
