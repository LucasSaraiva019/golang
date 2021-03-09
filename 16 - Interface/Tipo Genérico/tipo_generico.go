package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(10)
	generica("string")
	generica(true)
}
