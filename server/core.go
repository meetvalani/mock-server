package mockserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	logger "mockserver/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpResponse struct {
	ID                      int               `json:"id"`
	Method                  string            `json:"method"`
	Endpoint                string            `json:"endpoint"`
	ResponseCode            int               `json:"responseCode"`
	HTTPResponseContentType string            `json:"httpResponseContentType"`
	HTTPHeaders             map[string]string `json:"httpHeaders"`
	HTTPResponseBody        string            `json:"httpResponseBody"`
}

func (hR *httpResponse) getFromRequest(req *http.Request) error {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reqBody, &hR)
	if err != nil {
		return err
	}
	return nil
}

func (hR *httpResponse) getString() string {
	data, err := json.Marshal(*hR)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func commonReqParser(w http.ResponseWriter, req *http.Request) (*httpResponse, error) {
	obj := &httpResponse{}
	err := obj.getFromRequest(req)
	if err != nil {
		logger.Error.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
		return obj, err
	}
	return obj, nil
}

func add(w http.ResponseWriter, req *http.Request) {
	obj, err := commonReqParser(w, req)
	if err != nil {
		return
	}

	query := "insert into mock values (?,?,?,?,?,?,?)"
	headers, err := json.Marshal(obj.HTTPHeaders)
	if err != nil {
		fmt.Fprintf(w, "Some error in headers unmarshaling: "+err.Error())
		return
	}
	_, err = database.Execute(
		query,
		obj.ID,
		obj.Method,
		obj.Endpoint,
		obj.ResponseCode,
		obj.HTTPResponseContentType,
		string(headers),
		obj.HTTPResponseBody,
	)
	if err != nil {
		fmt.Fprintf(w, "Some error in db: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func get(w http.ResponseWriter, req *http.Request) {

	var objs []httpResponse
	var obj httpResponse
	query := `select
	id, method, endpoint, responseCode, httpResponseContentType, httpHeaders, httpResponseBody
	from mock`
	rows, err := database.Select(query)
	if err != nil {
		fmt.Fprintf(w, "Some error in db: "+err.Error())
	}
	defer rows.Close()

	empty := true
	for rows.Next() {
		empty = false
		var id, responseCode int
		var method, endpoint, httpResponseContentType, httpHeaders, httpResponseBody string
		rows.Scan(&id, &method, &endpoint, &responseCode, &httpResponseContentType, &httpHeaders, &httpResponseBody)
		var headers = make(map[string]string)
		json.Unmarshal([]byte(httpHeaders), &headers)
		obj = httpResponse{
			ID:                      id,
			Method:                  method,
			Endpoint:                endpoint,
			ResponseCode:            responseCode,
			HTTPResponseContentType: httpResponseContentType,
			HTTPHeaders:             headers,
			HTTPResponseBody:        httpResponseBody,
		}
		objs = append(objs, obj)
	}
	if empty {
		fmt.Fprintf(w, "{}")
		return
	}
	returnResponse, err := json.Marshal(objs)
	if err != nil {
		logger.Error.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(returnResponse))
}

func getByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idS := params["id"]
	id, _ := strconv.Atoi(idS)
	var obj httpResponse
	query := fmt.Sprintf(
		`select method, endpoint, responseCode, httpResponseContentType, httpHeaders, httpResponseBody 
		from mock where id=%d`,
		id,
	)
	rows, err := database.Select(query)
	if err != nil {
		fmt.Fprintf(w, "Some error in db: "+err.Error())
	}
	defer rows.Close()

	empty := true
	for rows.Next() {
		empty = false
		var responseCode int
		var method, endpoint, httpResponseContentType, httpHeaders, httpResponseBody string
		rows.Scan(&method, &endpoint, &responseCode, &httpResponseContentType, &httpHeaders, &httpResponseBody)
		var headers = make(map[string]string)
		json.Unmarshal([]byte(httpHeaders), &headers)
		obj = httpResponse{
			ID:                      id,
			Method:                  method,
			Endpoint:                endpoint,
			ResponseCode:            responseCode,
			HTTPResponseContentType: httpResponseContentType,
			HTTPHeaders:             headers,
			HTTPResponseBody:        httpResponseBody,
		}
	}
	if empty {
		fmt.Fprintf(w, "{}")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func update(w http.ResponseWriter, req *http.Request) {
	obj, err := commonReqParser(w, req)
	if err != nil {
		return
	}

	query := `update mock set 
	method = ?,
	endpoint = ?,
	responseCode = ?,
	httpResponseContentType = ?,
	httpHeaders = ?,
	httpResponseBody = ?
	where id=?`
	headers, err := json.Marshal(obj.HTTPHeaders)
	if err != nil {
		fmt.Fprintf(w, "Some error in headers marshling: "+err.Error())
		return
	}
	_, err = database.Execute(
		query,
		obj.Method,
		obj.Endpoint,
		obj.ResponseCode,
		obj.HTTPResponseContentType,
		string(headers),
		obj.HTTPResponseBody,
		obj.ID,
	)
	if err != nil {
		fmt.Fprintf(w, "Some error in db: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	ids := params["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		logger.Error.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}

	query := "DELETE FROM mock WHERE id=(?)"
	_, err = database.Execute(query, id)
	if err != nil {
		logger.Error.Printf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "deleted successfully if present."}`)
}
