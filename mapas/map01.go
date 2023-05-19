package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&str)

	fmt.Println(caracteres(str))
}

func caracteres(str string) map[string]int {
	ocorrencias := make(map[string]int)

	for i := range str {
		car := string(str[i])
		ocorrencias[car]++
	}

	return ocorrencias
}
