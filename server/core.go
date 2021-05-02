package mockserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Println(err.Error())
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
	// some db work
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func get(w http.ResponseWriter, req *http.Request) {
	obj, err := commonReqParser(w, req)
	if err != nil {
		return
	}
	// some db work
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func update(w http.ResponseWriter, req *http.Request) {
	obj, err := commonReqParser(w, req)
	if err != nil {
		return
	}
	// some db work
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}

func delete(w http.ResponseWriter, req *http.Request) {
	obj, err := commonReqParser(w, req)
	if err != nil {
		return
	}
	// some db work
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, obj.getString())
}
