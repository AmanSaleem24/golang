package main

import "fmt"

func main(){
	count:= 10
	for i := 0; i < count; i++ {
		fmt.Print(i, " ")
	}
	for{
		fmt.Print(count, " ")
	}
	for i:= range count {
		fmt.Print(i, "")
	}
}

