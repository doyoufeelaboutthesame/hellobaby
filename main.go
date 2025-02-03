package main

import (
	"FirstProject/internal/database"
	"FirstProject/internal/handlers"
	"FirstProject/internal/taskService"
	"FirstProject/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	database.Db.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.Db)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8000"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
