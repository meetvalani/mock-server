package mockserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	localDatabase "mockserver/database"
	logger "mockserver/logger"

	"github.com/gorilla/mux"
)

var database *localDatabase.Database

func welcomePage(w http.ResponseWriter, req *http.Request) {
	// TODO: return graphical UI.
	fmt.Fprintf(w, "This is under progress.\n")
}

func getResponseFromDB(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	path := params["path"]

	query := `select method, responseCode, httpResponseContentType, httpHeaders, httpResponseBody
	from mock where endpoint=(?)`
	statement, err := database.PreparedStatement(query)
	if err != nil {
		logger.Error("Error in prepare statement: ", err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	results, err := statement.Query(path)
	if err != nil {
		logger.Error("Error in select query: ", err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	empty := true
	var responseCode int
	var method, httpResponseContentType, httpHeaders, httpResponseBody string
	for results.Next() {
		empty = false
		results.Scan(&method, &responseCode, &httpResponseContentType, &httpHeaders, &httpResponseBody)
		break
	}
	results.Close()
	statement.Close()

	if empty {
		w.WriteHeader(404)
		fmt.Fprintf(w, "404 page not founddd\n")
		return
	}

	if req.Method != strings.ToUpper(method) {
		w.WriteHeader(405)
		return
	}

	headers := make(map[string]string)
	err = json.Unmarshal([]byte(httpHeaders), &headers)
	if err != nil {
		logger.Error("Error in returning headers: ", err)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", httpResponseContentType)

	for key, value := range headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(responseCode)
	fmt.Fprintf(w, httpResponseBody)
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
	logger.Info("Starting server at: '%s'", address)
	srv.ListenAndServe()
}
