package servidor

import (
	banco "crud/24.1-banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"ID"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requsição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
	}
	fmt.Println(usuario)

	db, erro := banco.Conctar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values(?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}

	IDinserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter ID inserido!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso ! ID: %d", IDinserido)))

}
