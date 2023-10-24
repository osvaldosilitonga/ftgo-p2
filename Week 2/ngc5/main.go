package main

import (
	"log"
	"net/http"
	"ngc5/config"
	"ngc5/handler"
	"ngc5/middleware"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.DBConn()
	defer db.Close()
	router := httprouter.New()

	userHandler := handler.NewUserHandler(db)
	adminHandler := handler.NewAdminHandler(db)

	// Routes
	router.POST("/login", middleware.Logging(userHandler.Login))
	router.POST("/register", middleware.Logging(userHandler.Register))
	router.GET("/admin", middleware.Logging(middleware.Auth(adminHandler.GetAdmin)))
	// router.POST("/register", middleware.Logging(middleware.BodyCheck(userHandler.Register)))

	log.Fatal(http.ListenAndServe(":8080", router))
}
