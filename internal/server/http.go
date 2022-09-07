package server

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type httpServer struct {
	log *Log
}

type httpAppendRequest struct {
	Data string `json:"data"`
}

type httpAppendResponse struct {
	Key string `json:"key"`
}

type httpGetResponse struct {
	Data string `json:"data"`
	Key string `json:"key"`
}

type httpGetAllReponse struct {
	Items []Record
}

func NewHTTPServer(addr string) (*http.Server) {
	server := &httpServer{
		log: NewLog(),
	}

	r := mux.NewRouter()
	r.HandleFunc("/", server.handleAppend).Methods("POST")
	r.HandleFunc("/", server.handleGetAll).Methods("GET")
	r.HandleFunc("/{key}", server.handleGet).Methods("GET")

	return &http.Server{
		Addr: addr,
		Handler: r,
	}
}

func (server *httpServer) handleAppend(w http.ResponseWriter, r *http.Request) {
	var request httpAppendRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, err := server.log.Append(request.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := httpAppendResponse{Key: key,}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (server *httpServer) handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	record, err := server.log.Get(vars["key"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := httpGetResponse{
		Key: record.Key,
		Data: record.Data,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (server *httpServer) handleGetAll(w http.ResponseWriter, r *http.Request) {
	records := server.log.GetAll()

	response := httpGetAllReponse{
		Items: records,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}