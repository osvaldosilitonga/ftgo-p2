package main

import (
	"avengers-v2/config"
	"avengers-v2/handler"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.DBConn()
	router := httprouter.New()

	inventoryHandler := handler.NewInventoryHandler(db)
	criminalHandler := handler.NewCriminalHandler(db)

	// Router
	router.GET("/heroes", handler.GetHeroes)
	router.GET("/villain", handler.GetVillain)
	// Inventory
	router.GET("/inventories", inventoryHandler.GetInventories)
	router.GET("/inventories/:id", inventoryHandler.GetInventoriesById)
	router.POST("/inventories", inventoryHandler.PostInventory)
	router.PUT("/inventories/:id", inventoryHandler.PutInventory)
	router.DELETE("/inventories/:id", inventoryHandler.DeleteInventory)
	// Criminal Reports
	router.GET("/criminal", criminalHandler.GetCriminalReports)
	router.GET("/criminal/:id", criminalHandler.GetCriminalReportById)
	router.POST("/criminal", criminalHandler.PostCriminalReport)
	router.PUT("/criminal/:id", criminalHandler.PutCriminalReport)
	router.DELETE("/criminal/:id", criminalHandler.DeleteCriminalReport)

	log.Fatal(http.ListenAndServe(":8080", router))
}
