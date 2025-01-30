package main

import (
	"FirstProject/internal/database"
	"FirstProject/internal/handlers"
	"FirstProject/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	database.Db.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.Db)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
