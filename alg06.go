package main

import "fmt"

func main() {
	var sal float64
	fmt.Print("Informe o salário atual: ")
	fmt.Scan(&sal)
	aum := sal + (sal / 100 * 15)
	fmt.Println("O novo salário com um aumento de 15% é igual a", aum)
}
