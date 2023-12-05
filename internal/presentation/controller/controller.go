package controller

import (
	"html/template"
	"mvc-golang/internal/domain"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := domain.GetAllUsers()
	var tmplFile = "/internal/presentation/controller/user_list.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	w.Header().Add("content-type", "text/html")
	err = tmpl.Execute(w, users)
	if err != nil {
		panic(err)
	}
}

func GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	announcements := domain.GetAllAnnouncements()
	var tmplFile = "/internal/presentation/controller/announcement_list.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	w.Header().Add("content-type", "text/html")
	err = tmpl.Execute(w, announcements)
	if err != nil {
		panic(err)
	}
}
