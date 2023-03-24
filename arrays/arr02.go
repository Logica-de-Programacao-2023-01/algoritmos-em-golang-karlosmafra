package main

import "fmt"

func main() {
	var nomes = []string{}
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	nomes = append(nomes, "José")
	fmt.Println(nomes)
}
