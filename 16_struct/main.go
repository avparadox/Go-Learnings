package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	phone string
}

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // It has a nano sec precision 
	customer
}

func (o *order) changeStatus (status string){
	o.status = status
}

// receiver type
func (a order) getAmount () float32 {
	return a.amount 
}

// Dynamic Instance creation of struct using Functions.
func newOrder  (id string, amount float32, status string) *order {
	myOrder := order{
		id : id,
		amount: amount,
		status: status,		
	}
	return &myOrder
}

func main (){

	// newCustomer:= customer{
	// 	name: "Adi",
	// 	phone: "1234",
	// }	

	newOrder := order{
		id: "23",
		amount: 45,
		status: "received",
		// customer:newCustomer,
		customer: customer{name: "Adi", phone: "124"},
	}

	fmt.Println(newOrder)
	newOrder.customer.name = "Aditya"
	fmt.Println(newOrder)
	
	// order := struct {
	// 	id string
	// 	amount float32
	// 	status string
	// }{"1", 100, "Done"}

	// fmt.Println(order)

	// langauge := struct{
	// 	name string
	// 	age float32
	// }{"Adi", 21}

	// fmt.Println(langauge)


	// myOrder :=  newOrder("1",22,"done")
	// // fmt.Println(newOrder("1",22, "Done"))
	// fmt.Println(myOrder.amount)


	// If you don't set any of the fields, then default values are enforced on the print statement for the specific variable/object.

	// Int -> 0
	// String -> Empty space " "
	
	// 	In Go, nil represents the zero value for several types, including:
	// Pointers: A pointer that doesn't point to anything is nil.
	// Slices: An empty slice is represented by nil.
	// Maps: An empty map is nil.
	// Channels: A channel that hasn't been initialized is nil.
	// Functions: A function that hasn't been assigned a value is nil.
	// Interfaces: An interface variable that doesn't hold any value is nil.

// 	Not all zero values are nil:
	// For example, the zero value of an integer is 0, not nil. Similarly, the zero value of a string is "" (an empty string), not nil.
// Comparing with nil:
	// You can use the == operator to check if a variable of the above types is nil


	// myOrder := order{
	// 	id : "1",
	// 	// amount: 100.00,
	// 	status: "Delivered",		
	// }
	
	// Phase 1 Started
	// Created sepeartely
	// myOrder.createdAt = time.Now()
	// println("")
	// fmt.Println("Total Order Details: ", myOrder)
	// println("")
	// println("Specific Details: ")
	// println("")
	// fmt.Println("Order id: ", myOrder.id)
	// fmt.Println("Order amount: ", myOrder.amount)
	// fmt.Println("Order status: ", myOrder.status)
	// fmt.Println("Order createdAt: ", myOrder.createdAt)
	// println("")
	// Phase 1 Ended


	// Phase 2 Started

	// myOrder2 := order{
	// 	id : "2",
	// 	amount: 23,
	// 	status: "In Progress",
	// 	createdAt: time.Now(),
	// }

	// myOrder.amount = 266;
	// fmt.Println(myOrder)
	// fmt.Println(myOrder2)

	// Phase 2 Ended

	// Phase 3
	// myOrder.changeStatus("Confirmed")
	// fmt.Println(myOrder)
	// fmt.Println(myOrder.status)
	// fmt.Println(myOrder.getAmount())
	// Phase 3 Ended
}