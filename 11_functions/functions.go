package main

import "fmt"

// func add (a, b int) int{
// 	return a + b
// }

// func languages() (string, string, int){
// 	return "golang", "javascript", 2
// }

func processIt(fn func(a int) int) (func(a int) int, int){
	d := fn(67)
	return func(a int) int {
		return a / 2
	}, d
}



func main () {

	fn := func (a int) int{
		return a*2
	}

	ret, d:= processIt(fn)

	fmt.Println(ret(d))

	// sum := add(3, 6)
	// fmt.Println(sum)

	// lang1, lang2, count := languages()

	// fmt.Println(lang1, lang2, count)
}


