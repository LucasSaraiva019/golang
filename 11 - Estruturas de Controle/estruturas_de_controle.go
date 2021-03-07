package main

import "fmt"

func main() {
	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	// A variável outroNumero só pode ser utilizada dentro do escopo (if;else)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	} else if numero < -10 {
		fmt.Println("Menor que 0")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
