package main

import "fmt"

// Passando a cópia do valor para a função
func inverterSinal(numero int) int {
	return numero * -1
}

// Passando a referência do ponteiro para a função
// Altera o código inteiro
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}
func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
