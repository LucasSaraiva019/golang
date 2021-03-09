package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(3)

	go func() {
		escrever("Olá mundo") // goroutine
		waitgroup.Done()      // Tira um cara do contator do Add(2)
	}()

	go func() {
		escrever("Programando em Go")
		waitgroup.Done()
	}()
	go func() {
		escrever("Goroutine 3")
		waitgroup.Done()
	}()
	waitgroup.Wait() // Espera a contagem das goroutines chegar a 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
