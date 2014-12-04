package controllers

import (
	"html/template"
	"net/http"
)

type App struct{}

func (a *App) Test() {

}

func (app *App) Index(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(rw, nil)
}
