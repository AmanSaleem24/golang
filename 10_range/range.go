package main

import "fmt"

func main(){

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	// for i := range(len(nums)) {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	for i, num:= range nums {
		sum += num
		fmt.Println(i, num, sum)
	}

	m:= map[string]string{"fname": "Aman", "lname": "Saleem"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for i, c := range "gOlAng" {
		fmt.Println(i, string(c))
	}
}