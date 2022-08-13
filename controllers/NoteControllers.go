package controllers

import (
	"fmt"
	"golang/models"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteControllers struct{}

func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var notes []models.Note
	db.Find(&notes)

	datas := map[string]interface{}{
		"Notes": notes,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
func (controller *NoteControllers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		note := models.Note{
			Assignee: r.FormValue("asignee"),
			Content:  r.FormValue("content"),
			Date:     r.FormValue("date"),
		}

		result := db.Create(&note)
		if result.Error != nil {
			log.Println(result.Error)
			fmt.Println(result.Error)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		files := []string{
			"./view/base.html",
			"./view/create.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "iternal server error", http.StatusInternalServerError)
			return
		}

		err = htmlTemplate.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "iternal server error", http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

}
