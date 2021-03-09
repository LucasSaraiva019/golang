package main

import (
	"fmt"
)

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	restoDaDivisao := 1 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores Unários
	numero := 10
	numero++
	numero += 15 // numero = nnumero + 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 20
	fmt.Println(numero)

	numero /= 20
	fmt.Println(numero)

}
