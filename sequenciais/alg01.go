package main

import "fmt"

func main() {
	var (
		n1 int
		n2 int
		n3 int
	)
	fmt.Print("Digite três números inteiros: ")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	soma := n1 + n2 + n3
	fmt.Println("A soma dos números é igual a", soma)
}
