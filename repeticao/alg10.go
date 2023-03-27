package main

import "fmt"

func main() {
	var x int
	var maior int
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	maior = x

	for x != 0 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		if x > maior {
			maior = x
		}
	}

	fmt.Println("O maior número é", maior)
}
