package main

import "fmt"

// num is not passed directly, a copy/vlaue of it is passed in the function runtime.
// func changeNum (num int) {
// 	num = 7
// 	fmt.Println("In Change Num: ", num)
// }

// By Reference

func changeNum(num *int){
	
	fmt.Println("Before In Change Function Call: ", &num)
	*num = 5
	fmt.Println("In Num Change Value", *num)
}

func main(){

	num := 10
	// changeNum(number)


	fmt.Println("Before Function Call: ", num)
	changeNum(&num)
	fmt.Println("After Function Num: ", num)

	// fmt.Println("Value of Number", number)
	// fmt.Println("Location of number in memory", &number)


}

