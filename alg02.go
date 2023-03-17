package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	fmt.Println("O dobro do número é igual a", num*2)
	fmt.Println("O triplo do número é igual a", num*3)
	fmt.Println("O quádruplo do número é igual a", num*4)
}
