package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
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
	//fmt.Fprintln(w, "В БД добавлены:")
	//fmt.Fprintln(w, "\"task\": ", message.Task)
	//fmt.Fprintln(w, "\"is_done\": ", message.IsDone)
	//fmt.Fprintln(w, "Полная информация:")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding json:", err)
		return
	}
}
func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)
	message := Message{}
	result := DB.First(&message, id)

	if result.Error != nil {
		http.Error(w, "There is no message on this ID", http.StatusInternalServerError)
		log.Println("Error getting message: ", result.Error)
		return
	}
	result = DB.Delete(&message)
	if result.Error != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		log.Println("Failed to delete message:", result.Error)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	log.Printf("Message with ID %d deleted successfully", id)
}
func PatchMessage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		log.Println("Error reading request body: ", err)
		return
	}
	defer req.Body.Close()
	var updateData Message
	if err := json.Unmarshal(body, &updateData); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		log.Println("Error decoding JSON: ", err)
		return
	}
	var message Message
	result := DB.First(&message, id)
	if result.Error != nil {
		http.Error(w, "There is no message on this ID", http.StatusNotFound)
		log.Println("Error getting message: ", result.Error)
		return
	}
	if updateData.Task != "" {
		message.Task = updateData.Task
	}
	if updateData.IsDone != message.IsDone {
		message.IsDone = updateData.IsDone
	}
	result = DB.Save(&message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Println("Error updating message: ", result.Error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding json:", err)
		return
	}
	log.Println("Successfully updated message")
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages/{id}", PatchMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
z