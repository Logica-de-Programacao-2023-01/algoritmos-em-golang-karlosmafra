package main

import "fmt"

func main() {
	var idd int
	fmt.Print("Qual a sua idade em anos? ")
	fmt.Scan(&idd)
	fmt.Println("A sua idade em dias é igual a", idd*365, "dias.")
}
