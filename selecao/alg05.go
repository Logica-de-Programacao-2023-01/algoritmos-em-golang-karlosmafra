package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e de 5 ao mesmo tempo")
	} else if num%3 == 0 {
		fmt.Println("O número é múltiplo de 3")
	} else if num%5 == 0 {
		fmt.Println("O número é múltiplo de 5")
	} else {
		fmt.Println("O número não é múltiplo de 3 nem de 5")
	}
}
