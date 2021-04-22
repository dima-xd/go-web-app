package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-web-app/pkg/data"
	"html/template"
	"log"
	"net/http"
)

type loginPage struct {
	data      *data.UserData
	templates *template.Template
}

func ServeLoginResource(r *mux.Router, data data.UserData, templates *template.Template) {
	api := &loginPage{data: &data, templates: templates}
	r.HandleFunc("/", api.loginPage).Methods("GET")
	r.HandleFunc("/submit", api.submitUser).Methods("POST")
}

func (a loginPage) loginPage(w http.ResponseWriter, r *http.Request) {
	a.templates.ExecuteTemplate(w, "login.html", nil)
}

func (a loginPage) submitUser(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	fmt.Println(login)
	isLoginExists, err := a.data.IsLoginExists(login)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	if isLoginExists {
		//there will be redirect to the next page
	} else {
		//there will be redirect to the page with errors
	}
}
