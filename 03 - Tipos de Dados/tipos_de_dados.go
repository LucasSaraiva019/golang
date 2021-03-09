package main

import (
	"errors"
	"fmt"
)

func main() {
	//Int
	// Usa a arquitetura do pc no meu caso int64
	var numero int64 = -10000
	fmt.Println(numero)

	//uint = unsygned int só positivo!
	var numero2 uint64 = 10000
	fmt.Println(numero2)

	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	// Strings
	var str string = "String"
	fmt.Println(str)

	str2 := "String2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// Valor Zero
	var valorZeroTexto string
	fmt.Println(valorZeroTexto)

	var valorZeroInt int
	fmt.Println(valorZeroInt)

	var valorZeroBool bool
	fmt.Println(valorZeroBool)

	var valorZeroErro error
	fmt.Println(valorZeroErro)
	//Boleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	// Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
