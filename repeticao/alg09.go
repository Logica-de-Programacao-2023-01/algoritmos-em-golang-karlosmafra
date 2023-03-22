package main

import "fmt"

func main() {
	var nums = []int{}
	fmt.Print("Digite números inteiros: ")
	//ler valores
	fmt.Scan(&nums)
	fmt.Println(nums)
	soma := 0
	for i := 0; i < len(nums); i++ {
		soma += nums[i]
		if nums[i] == 0 {
			break
		}
	}
	med := soma / len(nums)
	fmt.Println("A média é igual a", med)
}
