package main

import (
	"fmt"
	"time"
)

func main(){

	day:= 8

	// simple switch
	switch day {
	case 1:fmt.Println("Monday")
	case 2:fmt.Println("Tuesday")
	case 3:fmt.Println("Wednesday")
	case 4:fmt.Println("Thursday")
	case 5:fmt.Println("Friday")
	case 6:fmt.Println("Saturday")
	case 7:fmt.Println("Sunday")
	default:fmt.Println("Invalid day number")
	}

	//  multiple condition switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend, Go Enjoy!", time.Now().Weekday())
	default:
		fmt.Println("It's a Workday-", time.Now().Weekday())
	}

	// type switch
	whoAmI := func (i interface{}) {
		switch t:= i.(type){
		case int: fmt.Println("It's an integer")
		case string: fmt.Println("It's a string")
		case float32, float64: fmt.Println("It's a float")
		case bool: fmt.Println("It's a bool")
		default: fmt.Println("It's other type", t)
		}
	}

	whoAmI(87.90)

}
