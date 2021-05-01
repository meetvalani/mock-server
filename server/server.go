package mockserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func welcomePage(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "This is under progress.\n")
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", welcomePage)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	srv.ListenAndServe()
}
