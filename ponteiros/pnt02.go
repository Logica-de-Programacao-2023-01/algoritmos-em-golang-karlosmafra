package main

import "fmt"

func main() {
	b := false
	var ptr *bool = &b

	fmt.Println("Valor de b:", b)

	inverter(ptr)

	fmt.Println("Novo valor de b:", b)

}

func inverter(ptr *bool) {
	if *ptr == true {
		*ptr = false
	} else {
		*ptr = true
	}
}
