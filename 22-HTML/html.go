package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	//rota para home com assinatura padrão ResponseWrite e Request
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Theo", "theo@educare.com.br"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	//comando para levantar na porta determinada (servidor rodando na porta 5000)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
