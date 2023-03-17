package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Peso em quilos: ")
	fmt.Scan(&peso)
	fmt.Println("O peso em libras é igual a", peso*2.205, "lb")
}
