package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Avenida Perimetral")
	fmt.Println(TipoEndereco)
}
