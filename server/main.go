package main

import (
	"log"
	"server/app"
	"server/controller"
	"server/repository"
	"server/service"
	"server/ws"
)

func main() {
	dbConn, err := app.NewDatabase()

	if err != nil {
		log.Fatalf("Unable to initialize database connction: %s", err)
	}

	userRepo := repository.NewRepository(dbConn.GetDB())
	userService := service.NewService(userRepo)
	userHandler := controller.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	app.InitRouter(userHandler, wsHandler)
	app.Start("0.0.0.0:8080")
}
