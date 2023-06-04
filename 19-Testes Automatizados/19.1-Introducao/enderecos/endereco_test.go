package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "avenida paulista"

	tipoDeEnderecoEsperado := "rua"

	tipoDeEnderecoRecebido := TipoDoEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}
