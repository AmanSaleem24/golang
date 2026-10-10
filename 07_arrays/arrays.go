package main

import "fmt"
func main(){

	// default values
	// int -> 0, bool -> false, string -> ""

	var nums [4]int
	// fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums)

	for i:= range len(nums) {
		nums[i] = i
	}
	// fmt.Println(nums)

	// intialising with the elements while declaring
	arr:= [3]int{1, 3, 5}
	fmt.Println(arr)

	// 2d array
	arr1 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr1)

	//array list <-> slices
	//dynamic sizing, constant time access, memory optimisation
	

	
}
