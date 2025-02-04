package main

import (
	"FirstProject/internal/database"
	"FirstProject/internal/handlers"
	"FirstProject/internal/taskService"
	"FirstProject/internal/userService"
	"FirstProject/internal/web/tasks"
	"FirstProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	err := database.Db.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}

	repo := taskService.NewTaskRepository(database.Db)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	//===========================================================================
	uerr := database.Db.AutoMigrate(&userService.Users{})
	if uerr != nil {
		log.Fatal(uerr)
	}

	urepo := userService.NewUserRepository(database.Db)
	uservice := userService.NewUserService(urepo)

	uhandler := handlers.NewUserHandler(uservice)

	//===========================================================================
	ue := echo.New()

	ue.Use(middleware.Logger())
	ue.Use(middleware.Recover())
	ue.Use(middleware.Logger())
	ue.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(ue, strictHandler)

	ustrictHandler := users.NewStrictHandler(uhandler, nil)
	users.RegisterHandlers(ue, ustrictHandler)

	if err := ue.Start(":8000"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
