package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float32
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index )
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{ 
			Nome: "Camisa",
			Descricao: "Azul",
			Preco: 40,
			Quantidade: 3,
		},
		{
			"Tenis", 
			"Preto", 
			50, 
			5, 
		},
		{
			"Fone", 
			"Ok", 
			200,
			3,
		},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}