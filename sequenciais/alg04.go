package main

import "fmt"

func main() {
	var (
		n2 float64
		n3 float64
		n5 float64
	)
	fmt.Print("Digite um número com peso 2: ")
	fmt.Scan(&n2)
	fmt.Print("Digite um número com peso 3: ")
	fmt.Scan(&n3)
	fmt.Print("Digite um número com peso 5: ")
	fmt.Scan(&n5)
	mediap := (n2*2 + n3*3 + n5*5) / 10
	fmt.Println("A média ponderada é igual a", mediap)
}
