package main

import "fmt"

func main() {
	x := 10
	y := 25
	var px *int = &x
	var py *int = &y

	fmt.Println("Valor de x:", *px, "  Valor de y:", *py)

	swap(px, py)

	fmt.Printf("Novo valor de x: %d   Novo valor de y: %d", *px, *py)
}

func swap(px, py *int) {
	var ptrx = *px
	*px = *py
	*py = ptrx
}
