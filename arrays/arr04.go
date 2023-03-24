package main

import "fmt"

func main() {
	var nums = [6]float64{2, 4.5, 5, 10, 20, 25.5}
	var soma float64
	for i := 0; i < len(nums); i++ {
		soma += nums[i]
	}
	med := soma / float64(len(nums))
	fmt.Println("A média é igual a", med)
}
