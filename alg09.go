package main

import "fmt"

func main() {
	var val float64
	fmt.Print("Preço do produto: ")
	fmt.Scan(&val)
	fmt.Println("O preço do produto com desconto de 10% é igual a", val-(val/10))
}
