package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

//método salvar
func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuário %s no banco\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario := usuario{"Theo", 17}
	fmt.Println(usuario)

	usuario.salvar()

	retorno := usuario.maiorIdade()

	if retorno == true {
		fmt.Printf("Usuário %s é maior de idade\n", usuario.nome)
	} else {
		fmt.Printf("Usuário %s não é maior de idade\n", usuario.nome)

	}

	usuario.fazerAniversario()

	retorno2 := usuario.maiorIdade()

	if retorno2 == true {
		fmt.Printf("Usuário %s é maior de idade\n", usuario.nome)
	} else {
		fmt.Printf("Usuário %s não é maior de idade\n", usuario.nome)

	}

}
