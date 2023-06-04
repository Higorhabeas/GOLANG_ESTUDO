package enderecos

import "testing"

type cenarioDeTeste struct {
	endetecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"rua 1", "rua"},
		{"avenida 1", "avenida"},
		{"rua 2", "rua"},
		{"rodovia 3", "rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDoEndereco(cenario.endetecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

	/*enderecoParaTeste := "Avenida paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDoEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}*/

}
