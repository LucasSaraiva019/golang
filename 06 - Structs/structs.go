package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	var usuario1 usuario
	usuario1.nome = "Davi"
	usuario1.idade = 21
	fmt.Println(usuario1)

	endereco2 := endereco{"Rua dos Bobos", 0}
	usuario2 := usuario{"Lucas", 20, endereco2}
	fmt.Println(usuario2)

}
