package controllers

import (
	"html/template"
	"net/http"

	"web-server/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", produtos)
}