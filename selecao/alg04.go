package main

import "fmt"

func main() {
	var (
		peso float32
		alt  float32
		clas string
	)
	fmt.Print("Informe o peso em kg: ")
	fmt.Scan(&peso)
	fmt.Print("Informe a altura em metros: ")
	fmt.Scan(&alt)
	imc := peso / (alt * alt)

	if imc < 18.5 {
		clas = "Abaixo do peso"
	} else if imc < 25 {
		clas = "Dentro do peso"
	} else {
		clas = "Acima do peso"
	}

	fmt.Println("O IMC é igual a", imc)
	fmt.Println("A classificação é:", clas, "ideal.")
}
