package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	age int
	sex string
	phone int64
}

func newCustomer (name string, age int, sex string, phone int64) *customer {
	cust := customer {
		name: name,
		age: age,
		sex: sex,
		phone: phone,
	}
	return &cust
}

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time      // nanosecond precision
	customer
}



func (o *order) updateStatus (status string) {
	o.status = status
}

func (o order) getAmount () float32{
	return o.amount
}

func newOrder (id string, amount float32, status string) *order {

	createOrder := order {
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}
	return &createOrder
}

func main(){

	// myOrder := order {
	// 	id: "1",
	// 	amount: 1200.00,
	// 	status: "paid",
	// }
	// myOrder.updateStatus("unpaid")
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder)

	// cust := newCustomer("Aman", 21, "Male", 9555303228)
	// fmt.Println(cust)
	// order1 := newOrder("21", 213.2, "paid")
	// order1.customer = *cust
	// fmt.Println(*order1)
	
	// order1.customer = customer {
	// 	name: "Aman",
	// 	age: 22,
	// 	sex: "Male",
	// 	phone: 92378489892,
	// }

	// fmt.Println(order1)

	language := struct {
		name string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
	
}
