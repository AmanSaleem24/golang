package main

import "fmt"


const age = 21
var gender = "Male"

const ( 
	name = "Saleem"
	class = "B.Tech"
)
func main(){
	const name = "Saleem"
	// age = 19
	gender = "Man"
	class := "M.Tech"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(gender)
	fmt.Println(name, class)
}

