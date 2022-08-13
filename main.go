package main

import (
	"fmt"
	"golang/controllers"
	"golang/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// KONFIGURASI DATABASE SQLITE
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())

	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteController := &controllers.NoteControllers{}

	router := httprouter.New()

	router.GET("/", noteController.Index)
	router.GET("/create", noteController.Create)
	router.POST("/create", noteController.Create)

	port := ":8080"
	fmt.Println("Aplikasi berjalan di http://localhost:8080")

	log.Fatal(http.ListenAndServe(port, router))
}
