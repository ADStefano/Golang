package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
	   )


/*
Arquivos de teste tem que terminar com _test;

A função de teste tem que começar com Test seguido de uma letra maiúscula;

O parâmetro padrão é t *testing.T;

go test ./... roda todos os testes no diretório;

--cover para ver a porcentagem de testes bem sucedidos;

--coverprofile <nome-do-arquivo> gera um txt com as linhas que foram testadas;

Para rodar em paralelo utiliza-se o comando t.Parallel(), só roda quem tiver o comando no teste;

Para ler o arquivo estranho que é gerado pelo coverprofile utiliza-se o comando go tool cover --func=<nome-do-arquivo>;

E pra ver quais linhas do código foram testadas troca-se o func port html, gera um html com highlight das linhas testadas;

O . antes do import é o alias, tira a necessidade de escrever o nome do módulo antes do nome da função
*/
func TestEndereco(t *testing.T) {
	t.Parallel()
	enderecoParaTeste := "Rua Ângelo Di Stefano"
	enderecoEsperado := "Rua"
	enderecoRecebido := Endereco(enderecoParaTeste)

	if enderecoEsperado != enderecoRecebido{
		t.Errorf("O tipo de endereço recebido é diferente do esperado!!\nEsperado: %s\nRecebido: %s",enderecoEsperado, enderecoRecebido)
	}
}

type CenarioDeTeste struct{
	enderecoInserido string
	enderecoEsperado string
}

func TestEnderecoV2(t *testing.T){
	t.Parallel()

	cenariosDeTeste := []CenarioDeTeste{
		{"Rua Ângelo Di Stefano", "Rua"},
		{"Avenida Márcia Di Stefano", "Avenida"},
		{"Estrada Érico Pinton", "Estrada"},
		{"Rodovia Biscoito e Bolacha", "Rodovia"},
		{"Rotatória Da Puta que Pariu","Endereço inválido"},
		{"","Endereço inválido"},
	}

	for _, cenario := range cenariosDeTeste{
		tipoDeEnderecoRecebido := Endereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.enderecoEsperado{
			t.Errorf("O tipo recebido %s é diferente do tipo esperado %s",
		tipoDeEnderecoRecebido, cenario.enderecoEsperado)
		}
	}
}