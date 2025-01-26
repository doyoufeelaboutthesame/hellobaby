package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHendler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloHendler).Methods("GET")
	http.ListenAndServe(":8000", router)
}
