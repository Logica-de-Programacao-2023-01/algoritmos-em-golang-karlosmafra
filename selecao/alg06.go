package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("Digite dois números: ")
	fmt.Scan(&n1, &n2)
	if n1 >= 0 && n2 >= 0 {
		fmt.Println("A multiplicação dos dois números é igual a", n1*n2)
	} else {
		fmt.Println("A soma dos dois números é igual a", n1+n2)
	}
}
