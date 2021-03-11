package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		//{"Praça ABC", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"RUA DOS ABC", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		//{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("o tipo recebido é diferente do esperado! esperava %s e recebeu %s", retornoRecebido, cenario.retornoEsperado)
		}
	}

}
func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste Quebrou")
	}
}
