package main

import (
	"fmt"
	"maps"
)

func main(){

	// m:= make(map[string]string)
	m:= map[string]int{"age": 21, "phone": 9555}

	//setting an example
	//key does not exist then returns zeroed value(default vallue of each data type)

	// m["name"] = "golang"
	fmt.Println(m["phone"])
	fmt.Println(len(m))

	// delete(m, "name")
	// clear(m)

	k, ok := m["age"]

	if ok {
		fmt.Println("All ok", k)
	} else {
		fmt.Println("Not ok")
	}

	fmt.Println(len(m))

	m1 := map[string]int{"age": 21, "phone": 9555}
	m2 := map[string]int{"age": 21, "phone": 9555}

	fmt.Println(maps.Equal(m1, m2)) 
}
