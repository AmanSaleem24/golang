package main

import "fmt"

func main(){
	name:= "Aman"
	age:= 14
	if age >= 18 {
		fmt.Println(name, "is", age, "years old and he is an adult")
	} else {
		fmt.Println(name, "is", age, "years old and he is not an adult")
	}
}
