package main

import "fmt"

func main() {
	var num int
	var dia string
	fmt.Print("Digite um número inteiro entre 1 e 7: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		dia = "Domingo"
	case 2:
		dia = "Segunda-feira"
	case 3:
		dia = "Terça-feira"
	case 4:
		dia = "Quarta-feira"
	case 5:
		dia = "Quinta-feira"
	case 6:
		dia = "Sexta-feira"
	case 7:
		dia = "Sábado"
	default:
		dia = "[Número Inválido]"
	}
	fmt.Println("O dia da semana correspondente é", dia)
}
