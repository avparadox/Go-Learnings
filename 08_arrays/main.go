package main

import "fmt" 


func main(){

	// Zero Values Init mai
	// String => "", Int => 0, Boolean => false

	var nums [4]int 
	
	// println(len(nums))
	nums[1] = 25;
	
	println(nums[1])
	
	// println(nums) // Gives error
	fmt.Println(nums) // Works because of fmt lib

	// False Values Init mai
	var vals[4]bool
	fmt.Println(vals)
	vals[2] = true;
	fmt.Println(vals)


	// Strings
	var names[3]string
	fmt.Println(names)
	names[0] = "golang"
	fmt.Println(names)

	// Adding elements while declaration
	number:=[3]int{1,2,3}
	fmt.Println(number)


	// 2D Arrays
	numbers := [2][2]int{{1,2},{3,4}}
	fmt.Println(numbers)


	// Usage:
	// - fixed size arrays only
	// - memory optimization
	// - constant time access
}