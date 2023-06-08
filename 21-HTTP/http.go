package main

import (
	"log"
	"net/http"
)

func main() {

	/* cliente faz a requisição - servidor processa a requisição e envia a resposta
	Request - Response

	as rotas são os caminhos : URI - Identificador do Recurso(/home,/usuarios...)
								método - GET (busca dados), POST(cadastrar dados), PUT(atualizar dados), DELETE(apagar)
	*/

	//rota para home com assinatura padrão ResponseWrite e Request
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	//rota para usuários com assinatura padrão ResponseWrite e Request
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando página de usuários!!!!"))
	})

	//comando para levantar na porta determinada (servidor rodando na porta 5000)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
