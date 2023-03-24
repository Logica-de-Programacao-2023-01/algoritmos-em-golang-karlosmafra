package main

import "fmt"

func main() {
	var soma, quant, x int
	fmt.Print("Digite um número ")
	fmt.Scan(&x)

	soma += x
	quant++

	for x != 0 {
		fmt.Print("Digite um número")
		fmt.Scan(&x)
		if x != 0 {
			soma += x
			quant++
		}
	}
	med := soma / quant

	fmt.Println("A média é igual a", med)
}
