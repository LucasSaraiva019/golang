package main

import "fmt"

// Pode causar Stackoverflow = estou de pilha!!
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(45)
	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
