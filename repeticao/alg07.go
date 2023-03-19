package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if n%3 == 0 {
			fmt.Print("Fizz ")
		} else if n%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(n, " ")
		}
	}
}
