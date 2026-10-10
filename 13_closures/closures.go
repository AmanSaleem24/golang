package main

import "fmt"

func counter() func() int {
	var count int = 2
	return func () int{
		count++
		return count
	}
}
func main(){
	fn := counter()
	fmt.Println(fn())
	fmt.Println(fn())
}
