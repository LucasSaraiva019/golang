package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "Sabado"
	default:
		return "Numero Invalido!"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quinta"
	case numero == 5:
		diaDaSemana = "Sexta"
	case numero == 6:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero Invalido!"
	}
	return diaDaSemana
}

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)
	fmt.Println("---------------------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
