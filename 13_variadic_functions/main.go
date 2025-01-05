package main

import "fmt"

func sum (nums ...int) int {
	total := 0;

	for _,num := range nums {
		total = total + num
	}
	return total
}


func main(){
	nums := []int{3,3,3,3,3}
	result := sum(nums...)
	fmt.Println(result)
}