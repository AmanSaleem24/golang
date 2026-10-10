package main

import "fmt"

type OrderStatus int

const (
	Unpaid OrderStatus = iota
	Paid
	Prepared
	Delivered
	Returned
	Refunded
)

func changeStatus(status OrderStatus){
	fmt.Println("changing order status to", status)
}

func main(){
	changeStatus(Unpaid)
}
