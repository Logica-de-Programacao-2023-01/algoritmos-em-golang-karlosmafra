package main

import "fmt"

func main() {
	var dias float64
	var valor float64
	fmt.Print("Informe o número de dias trabalhados: ")
	fmt.Scan(&dias)
	fmt.Print("Informe o valor da diária: ")
	fmt.Scan(&valor)
	fmt.Println("O salário é igual a R$", dias*valor)
}
