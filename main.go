package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		 panic(err.Error())
	}

	return db
}

type Produto struct {
	id int
	Nome string
	Descricao string
	Preco float32
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close() // fecha conexão | defer é a ultima coisa executada na função, ele esperará tudo acabar antes de ser ativado
	http.HandleFunc("/", index )
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDosProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDosProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float32

		err = selectDosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)


	}
	 

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}