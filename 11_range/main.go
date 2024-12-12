package main

import (
	"fmt"
)

func main()  {
	 // Iterating over data structures.

	 nums := []int{5,6,7,8}
	// Normal loop
	//  for i:= 0; i <len(nums); i++{
	// 	fmt.Println(nums[i])
	//  }

	// Range
	for i, num := range nums{
		fmt.Println(i, num)
	}

	// m:= map[string]string{"fname": "john", "lname": "doe"}

	// Range on Maps
	// for k,v := range m{
	// 	fmt.Println(k,v)
	// }

	// Unicode | code point rune
	// 0,1,2 it is not index. It is the start of Rune.
	for i,c := range "golang"{
		fmt.Println(i,string(c))
	}
}