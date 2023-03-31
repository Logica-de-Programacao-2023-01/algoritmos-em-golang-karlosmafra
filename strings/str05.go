package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	fmt.Print(str)

}
