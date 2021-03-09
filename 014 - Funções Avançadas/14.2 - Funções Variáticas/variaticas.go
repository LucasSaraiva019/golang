package main

import "fmt"

// O numero variático só pode esta no final!
// Só pdoe ter um numero variático na função!
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma()
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
