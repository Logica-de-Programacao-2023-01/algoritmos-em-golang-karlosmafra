package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n <= len(str) && n >= 0 {
		fmt.Println(strings.ToUpper(str[0:n]) + str[n:])
	} else {
		fmt.Println("Número inválido")
	}
}
