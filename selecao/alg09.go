package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		num3 float64
	)
	fmt.Print("Digite três números reais: ")
	fmt.Scan(&num1, &num2, &num3)
	if num1 <= num2 && num1 <= num3 {
		if num2 <= num3 {
			fmt.Println("A sequência em ordem crescente é:", num1, num2, num3)
		} else {
			fmt.Println("A sequência em ordem crescente é:", num1, num3, num2)
		}
	} else if num2 <= num1 && num2 <= num3 {
		if num1 <= num3 {
			fmt.Println("A sequência em ordem crescente é:", num2, num1, num3)
		} else {
			fmt.Println("A sequência em ordem crescente é:", num2, num3, num1)
		}
	} else {
		if num1 <= num2 {
			fmt.Println("A sequência em ordem crescente é:", num3, num1, num2)
		} else {
			fmt.Println("A sequência em ordem crescente é:", num3, num2, num1)
		}
	}
}
