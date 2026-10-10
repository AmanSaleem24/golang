package main

import "fmt"

// func changeNum(num int) int {
// 	num = 3
// 	fmt.Println(num)
// 	return num
// }

func changeRefNum(a *int) {
	fmt.Println(a)
	*a = 5
	fmt.Println(*a)
}

func main(){
	num := 1

	fmt.Println(num)
	// changeNum(num)
	changeRefNum(&num)

	fmt.Println(num)
}

