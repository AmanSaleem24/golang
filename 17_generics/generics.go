package main

import (
	"fmt"
)

func printSlice[T any](nums []T){
	for _, item:= range nums {
		fmt.Print(item, " ")
	}
}

type Stack[T int | string | bool] struct {    // [T comparables	]
	list []T
}

func main(){
	stack := Stack[int] {
		list: []int{1, 2, 3, 4},
	}
	fmt.Println(stack)
	nums :=  []string{"golang"}
	printSlice(nums)
}


