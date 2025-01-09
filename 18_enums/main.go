package main

import "fmt"

// enumrated types
type orderStatus string 
type orderStatusNum int
const (
	Received orderStatus = "recevied"
	Confirmed 			 = "confirmed" 
	Prepared			 = "prepared"
	Delivered		 	 = "delivered"
)

const (
	Received1 orderStatusNum = iota
	Confirmed1 			  
	Prepared1			 
	Delivered1		 	 
)

func changeOrderStatus (status orderStatus) {
	fmt.Println("Changing order status to", status)
}

func changeOrderStatusNum (status orderStatusNum) {
	fmt.Println("Changing order status to", status)
}

func main(){
	changeOrderStatus(Confirmed)
	changeOrderStatusNum(Prepared1)
}