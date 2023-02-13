package enderecos

import (
	"strings"
)

//Comentário deve começar com o nome da função;
// Endereco: Verifica se o endereço tem um tipo válido e o retorna
func Endereco(enderecos string) string{
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	endereco := strings.ToLower(enderecos)
	primeira_palavra := strings.Split(endereco, " ")[0]
	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeira_palavra{
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeira_palavra)
	}

	return "Endereço inválido"
}