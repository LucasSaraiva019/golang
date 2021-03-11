package enderecos

import "strings"

// TipoDeEndereco Verifica se um endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	// Converte endereco para minuscula
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	// RUA TESTE = ["RUA", "TESTE"] -> Ira retornar o "RUA"
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}
	return "Tipo Inválido"

}
