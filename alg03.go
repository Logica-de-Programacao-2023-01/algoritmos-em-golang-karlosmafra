package main

import "fmt"

func main() {
	var peso float64
	var alt float64
	fmt.Print("Informe o peso da criança em kg: ")
	fmt.Scan(&peso)
	fmt.Print("Informe a altura da criança em metros: ")
	fmt.Scan(&alt)
	imc := peso / (alt * alt)
	fmt.Print("O IMC da criança é igual a ", imc)
}
