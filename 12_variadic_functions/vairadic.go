package main

import "fmt"

func sum(nums ...int) int{
	summation := 0
	for _, v := range nums {
		summation += v
	}
	return summation
}

func main(){
	fmt.Println(1, 2, 3, 4, 5, "hello", true, 12.3)
	nums := []int{1, 2, 32, 2, 5, 12, 756, 12}
	summation := sum(1, 2, 32, 2, 5, 12, 756, 12)
	fmt.Println(summation)
	total := sum(nums...)
	fmt.Println(total)
}
