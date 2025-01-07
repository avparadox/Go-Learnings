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

	fmt.Println("Order Struct: ", myOrder)

}