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

	_, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

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

func (controller *NoteControllers) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/base.html",
	}

	templateHtml, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var note models.Note
	db.Find("ID = ?", params.ByName("id")).Find(&note)

	data := map[string]interface{}{
		"Note": note,
		"ID":   params.ByName("id"),
	}

	err = templateHtml.ExecuteTemplate(w, "base", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}

}

func (controller *NoteControllers) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// noteID, _ := strconv.Atoi(params.ByName("id"))
	noteID := params.ByName("id")

	var note models.Note
	db.Where("ID = ?", noteID).First(&note)

	note.Assignee = r.FormValue("assignee")
	note.Date = r.FormValue("date")
	note.Content = r.FormValue("content")

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)

}

func (controller *NoteControllers) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
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
}

func (controller *NoteControllers) Done(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var note models.Note

	db.Find(&note, params.ByName("id"))

	note.IsDone = !note.IsDone

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)

}

func (controller *NoteControllers) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	var note models.Note

	db.Delete(&note, params.ByName("id"))

	http.Redirect(w, r, "/", http.StatusFound)

}
