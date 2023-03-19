package main

import "fmt"

func main() {
	var (
		n1    int
		n2    int
		n3    int
		menor int
	)
	fmt.Print("Digite três números inteiros: ")
	fmt.Scan(&n1, &n2, &n3)
	menor = n1
	if n2 < n1 {
		menor = n2
	}
	if n3 < menor {
		menor = n3
	}
	fmt.Println("O menor valor é", menor)
}
