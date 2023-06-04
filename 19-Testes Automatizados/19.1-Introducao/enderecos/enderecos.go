package enderecos

import "strings"

//TipoDoEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "esrtrada", "rodovia"}

	enderecoemletraminuscula := strings.ToLower(endereco)

	primeiraPalavradoEndereco := strings.Split(enderecoemletraminuscula, " ")[0]

	enderecotemtipovalido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavradoEndereco {
			enderecotemtipovalido = true
		}
	}

	if enderecotemtipovalido {
		return primeiraPalavradoEndereco
	}
	return "Tipo de enddereço inválido!"

}
