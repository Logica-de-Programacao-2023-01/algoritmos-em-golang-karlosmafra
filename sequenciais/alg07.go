package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	fmt.Println("O antecessor do número é", num-1)
	fmt.Println("O sucessor do número é", num+1)
}
