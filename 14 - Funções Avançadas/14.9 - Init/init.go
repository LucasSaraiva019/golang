package main

import "fmt"

var n int

// Ela é sempre executada antes da funnção Main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
