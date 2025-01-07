package main

import (
	"fmt"
	"time"
)

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // It has a nano sec precision 
}

func main (){

	
	myOrder := order{
		id : "1",
		amount: 100.00,
		status: "Delivered",		
	}
	
	// Phase 1 Started


	// Created sepeartely
	myOrder.createdAt = time.Now()
	println("")
	fmt.Println("Total Order Details: ", myOrder)
	println("")
	println("Specific Details: ")
	println("")
	fmt.Println("Order id: ", myOrder.id)
	fmt.Println("Order amount: ", myOrder.amount)
	fmt.Println("Order status: ", myOrder.status)
	fmt.Println("Order createdAt: ", myOrder.createdAt)
	println("")

	// Phase 1 Ended

	// Phase 2 Started

	myOrder2 := order{
		id : "2",
		amount: 23,
		status: "In Progress",
		createdAt: time.Now(),
	}

	myOrder.amount = 266;
	fmt.Println(myOrder)
	fmt.Println(myOrder2)

}