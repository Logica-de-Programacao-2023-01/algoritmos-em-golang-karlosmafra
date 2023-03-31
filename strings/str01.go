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
	str1 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	str2 := scanner.Text()

	str := str1 + " " + str2
	fmt.Println(str)
}
