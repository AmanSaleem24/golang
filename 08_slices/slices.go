package main

import "fmt"

func main(){

	//array list <-> slices
	//dynamic sizing, constant time access, memory optimisation
	//most used construct + usefult methods 

	//uninitialised slice is nil(null)
	
	// var nums []int 
	// fmt.Println(nums)			    		// []	
	// fmt.Println(nums == nil)			    	// true

	// //initialised slice declaration
	// var list = make([]int, 0, 5)
	// fmt.Println(list)                       // [0, 0]
	// fmt.Println(list == nil)					// false

	// //cap -> max no. of elements that can fit in the list at the moment
	// fmt.Println("cap", cap(list))					// 2, 5
	// list = append(list, 1)
	// list = append(list, 2)
	// list = append(list, 3)
	// list = append(list, 4)
	// fmt.Println("cap", cap(list))

	// // another way to initialise the slice
	// num := []int{}
	// fmt.Println("cap", cap(num))					// 2, 5
	// num = append(num, 1)
	// num[0] = 2
	// fmt.Println("cap", cap(num))
	
	// //copy function
	
	// num2 := make([]int, len(num))
	// copy(num2, num)
	// fmt.Println(num, num2)


	// Slicing
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums[1:])
}
