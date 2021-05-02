package mockserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
func StartServer(address string) {
	r := mux.NewRouter()
	r.HandleFunc("/", welcomePage).Methods("GET")
	r.HandleFunc("/get", get).Methods("GET")
	r.HandleFunc("/add", add).Methods("POST")
	r.HandleFunc("/update", update).Methods("PUT")
	r.HandleFunc("/delete", delete).Methods("DELETE")
	r.HandleFunc("/{path:[\\w\\W]+}", getResponseFromDB)

	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}
	log.Printf("Starting server at: '%s'", address)
	srv.ListenAndServe()
}
