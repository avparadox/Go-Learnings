package main

import "fmt" 


func main(){

	// Zero Values Init mai
	// String => "", Int => 0, Boolean => false

	var nums [4]int 
	
	// println(len(nums))
	nums[1] = 25;
	nums[2] = 255;
	
	// println(nums[1])
	// println(nums[2])
	
	// fmt.Println(len(nums))
	// println(nums) // Gives error
	// fmt.Println(nums) // Works because of fmt lib

	// False Values Init mai
	var vals[4]bool
	// fmt.Println(vals)
	vals[2] = true;
	// fmt.Println(vals)


	// Strings
	var names[3]string
	// fmt.Println(names)
	names[0] = "golang"
	// 1st position is being skipped and not showed like Int or Bool.
	names[2] = "Aditya"
	// fmt.Println(names)
	// Space is reserved but not being used and shadow is being returned.
	// fmt.Println(len(names))

	// Adding elements while declaration
	// number:=[3]int{1,2,3}
	// fmt.Println(number)

	// var name -> size of the arr -> type of arr > {values} -> cool hai
	// num2 :=[4]int{4,56,6}
	// fmt.Println(num2)

	// // 2D Arrays
	numbers := [2][2]int{{1,2},{3,4}}
	fmt.Println(numbers) 

	// 3D Arrays -> 3 times [2] means it is a 3D array and each array can have only 2 values 0th and 1st position. Play with it, then u can get it better.
	num2 := [2][2][2]int{{{1,2},{1,3}},{{1,4},{2,4}}}
	fmt.Println(num2)

	// Usage:
	// - fixed size arrays only
	// - memory optimization
	// - constant time access
}