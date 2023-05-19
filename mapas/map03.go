package main

import "fmt"

func main() {
	var slc = []int{2, 6, 3, 7, 9, 2, 6, 11, 8, 3, 9}

	fmt.Println("Os valores únicos são:", unicos(slc))
}

func unicos(slc []int) []int {
	mapa := make(map[int]int)
	var slcUnicos = []int{}

	for _, valor := range slc {
		mapa[valor]++
	}

	for chave := range mapa {
		if mapa[chave] == 1 {
			slcUnicos = append(slcUnicos, chave)
		}
	}

	return slcUnicos
}
