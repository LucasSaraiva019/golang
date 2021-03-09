package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	// For init = o j só pode ser usado dentro do for
	for j := 0; j < 10; j += 5 {
		println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"João", "Davi", "Lucas"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Nunca vou sair daqui")
		time.Sleep(time.Second)
	}
}
