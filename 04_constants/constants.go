package main

import "fmt"

const age = 23
const name = "aditya"


func main(){
	const name = "Aditya"
	// name = "Adi"  (Will give an err as constants cannot be changed)
	fmt.Println(name)
	fmt.Println(age)	

	const (
		short_name = "adi"
		age = 21
	)
	fmt.Println(short_name)
}