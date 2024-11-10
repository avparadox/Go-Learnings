package main

import (
	"time"
)


func main(){
	i := 3

	// Normal Switch
	switch i {
	case 1:
		println("Value of i is 1")
	case 2:
		println("Value of i is 2")
	case 3:
		println("Value of i is 3")
	default:
		println("Value of i is more than 3 or less than 1")
	}

	// Multiple Condition Switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		println("It is Weekend!")
	default:
		println("Kaam kar le bhai")
	}

	// Type Switch
	whoAmI := func (i interface{})  {
		switch i.(type){
		case int:
			println("It is an interger")		
		case string:
			println("it is a string")
		case bool:
			println("It is a Boolean")
		default:
			println("It is a type of others")
	}
}

whoAmI("Aditya")
}