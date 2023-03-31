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
	frase := scanner.Text()

	var car string
	fmt.Print("Digite um caractere: ")
	fmt.Scan(&car)

	var cont int
	for i := range frase {
		var letra = string(frase[i])
		if strings.Contains(letra, car) {
			cont++
		}
		i++
	}
	fmt.Printf("O caractere '%s' aparece %d vezes", car, cont)
}
