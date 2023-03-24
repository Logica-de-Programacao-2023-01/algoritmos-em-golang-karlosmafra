package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	nums = append(nums[:2], nums[3:]...)
	fmt.Println(nums)
}
