package servidor

import (
	banco "crud/24.1-banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		return
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

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conctar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de daos!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)

	}
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuário em JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)

	ID, erro := strconv.ParseUint(paramentros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro em inteiro"))
		return
	}
	db, erro := banco.Conctar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id= ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar id do usuário."))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scanear o usuário"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário em JSON"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	paramtros := mux.Vars(r)

	ID, erro := strconv.ParseUint(paramtros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro em inteiro."))
		return
	}

	corpRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler a requisição."))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter converter usuário em struct."))
		return
	}

	db, erro := banco.Conctar()
	if erro != nil {
		w.Write([]byte("Erro ao conctar ao banco de dados"))
		return
	}
	defer db.Close()

	statment, erro := db.Prepare("update usuarios set nome=?,email=? where id=?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statment.Close()

	if _, erro := statment.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar usuário."))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro id em inteiro."))
		return
	}

	db, erro := banco.Conctar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco."))
		return
	}
	defer db.Close()
	statement, erro := db.Prepare("delete from usuarios where id= ?")
	if erro != nil {
		w.Write([]byte("Erro ao preparar statement."))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuário."))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
