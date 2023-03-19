package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	for n := 1; n <= 10; n++ {
		fmt.Println(num, "x", n, "=", num*n)
	}
}
