package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

var message string

func HelloHendler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", message)
}
func PostHendler(w http.ResponseWriter, req *http.Request) {

	var r requestBody

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message = r.Message
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Message: %s", message)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/messagep", PostHendler).Methods("POST")
	router.HandleFunc("/messageg", HelloHendler).Methods("GET")
	http.ListenAndServe(":8000", router)
}
