package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func GetMessage(w http.ResponseWriter, req *http.Request) {
	var messages []Message
	result := DB.Find(&messages)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Println("Error getting messages from DB:", result.Error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding json:", err)
		return
	}
}
func CreateMessage(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		log.Println("Error reading request body: ", err)
		return
	}
	defer req.Body.Close()
	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		log.Println("Error decoding JSON: ", err)
		return
	}
	log.Println("Received: ", message)
	fmt.Println("Task: ", message.Task)
	fmt.Println("IsDone: ", message.IsDone)

	result := DB.Create(&message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Println("Error creating data in DB: ", result.Error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Data created successfully!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	http.ListenAndServe(":8080", router)
}
