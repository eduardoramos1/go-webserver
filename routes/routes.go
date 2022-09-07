package routes

import (
	"net/http"
	"web-server/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}