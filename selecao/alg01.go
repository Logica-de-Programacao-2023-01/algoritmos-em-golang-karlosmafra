package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("Digite dois números inteiros: ")
	fmt.Scan(&n1, &n2)
	if n1 > n2 {
		fmt.Println("O maior número é", n1)
	} else {
		fmt.Println("O maior número é", n2)
	}
}
