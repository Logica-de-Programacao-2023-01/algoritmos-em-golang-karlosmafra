package main

import "fmt"

func main() {
	var sal float64
	fmt.Print("Salário atual: R$")
	fmt.Scan(&sal)
	if sal < 1000 {
		fmt.Println("O salário com aumento é igual a R$", sal+sal/10)
	} else {
		fmt.Println("O salário com aumento é igual a R$", sal+(sal/100*5))
	}
}
