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

}