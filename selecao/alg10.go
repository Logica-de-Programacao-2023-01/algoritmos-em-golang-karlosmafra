package main

import "fmt"

func main() {
	var idd int
	var cls string
	fmt.Print("Informe a idade: ")
	fmt.Scan(&idd)
	if idd <= 9 {
		cls = "Mirim"
	} else if idd <= 13 {
		cls = "Infantil"
	} else if idd <= 17 {
		cls = "Juvenil"
	} else {
		cls = "Adulto"
	}
	fmt.Println("A classificação é:", cls)
}
