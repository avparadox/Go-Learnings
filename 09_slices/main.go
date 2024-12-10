package main

import (
	// "fmt"
	// "slices"
)

// Slices => Dynamic Arrays
// useful methods
func main(){

	// uninit slices are nil === null
	// Array is defined, no init of input values -> {}.

	var nums[]int  // -> Start Point
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))	
	nums = append(nums, 1,3)
	// fmt.Println(nums)


	// Make Method
	// Not a nil size
	// Making a array bit

	// var nums = make([]int, 0) 
	
	// var nums2 = make([]int, 3,5) 
	// Capacity is 5 and 3 signifies 3 times zero(0) in the array -> [0 0 0]
	// fmt.Println(cap(nums2)) // -> Answer is 5 (Capacity)
	
	// var nums3 = make([]int, 0,5)
	// nums3 = append(nums3, 2,5)
	// fmt.Println(nums3)
	
	// Cap means capacity -> maximum number of elements can fit.
	
	// Adding elements from the end
	// nums = append(nums, 2,1,2,3,4,5,6,7) 

	// fmt.Println(nums)
	// fmt.Println(len(nums)) // Cap is 8 -> total elements present in the array
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// Array [] and input {} is defined.
	numbers := []int{}
	// when u define a slice, it is not nil. It takes some memory in the system, so it will not be nil.

	// fmt.Println(numbers == nil)
	// fmt.Println(numbers)
	// fmt.Println(cap(numbers)) // -> 0
	// fmt.Println(len(numbers)) // -> 0

	numbers = append(numbers, 1,2,3,4)
	// fmt.Println(len(numbers))

	numbers = append(numbers, 1,2,3,4)
	// fmt.Println(len(numbers))
	
	numbers = append(numbers, 1,2,3,4)
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))
	// fmt.Println(cap(numbers))

	// Moral -> Capacity is doubled if needed and if the array fits well, it doesn't change the size of the array. Cool hai.

	// var temp = make([]int,1,5)
	// temp = append(temp, 1,4)
	// fmt.Println(len(temp))
	// fmt.Println(cap(temp))
	// temp[0] = 3
	// fmt.Println(temp)

	// Copy Function
	var copy1 = make([]int, 0,5)
	copy1 = append(copy1, 2)
	var copy2 = make([]int, len(copy1))

	// copy
	copy(copy2, copy1)
	// fmt.Println(copy1, copy2)

	// Slice Operator
	// 0:2 -> means start from 0th Index and go upto to 2nd Index and exclude the last value[2] here.
	// var slice = []int{0,1,2}
	// fmt.Println(slice[0:2])
	
	// Start from first
	// fmt.Println(slice[:2])
	
	// Go till last of the array
	// fmt.Println(slice[0:])

	// Comparing of slices
	// var slice1 = []int{1,2}
	// var slice2 = []int{1,2}
	
	// returns a bool
	// fmt.Println(slices.Equal(slice1,slice2))

	// 2D Arrays in Slices
	// var slice3 = [][]int{{1,2,3},{4,5,6}}
	// fmt.Println(slice3)
}	