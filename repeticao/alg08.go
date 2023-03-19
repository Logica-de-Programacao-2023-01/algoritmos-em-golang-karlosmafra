package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)
	fmt.Println("Divisores:")
	for d := 1; d <= num; d++ {
		if num%d == 0 {
			fmt.Print(d, " ")
		}
	}
}
