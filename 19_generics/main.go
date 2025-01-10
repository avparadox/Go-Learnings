package main

import "fmt"

func printSlice[T comparable, N string](items []T, name N){
	for _, item := range items{
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string){
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

// Stack -> LIFO
type Stack[T string | int] struct{
	elements [] T
}


func main(){

	myStack := Stack[int]{
		elements: []int{2,3,45,56,7,8,8,9},
	}

	fmt.Println(myStack)



	// nums := []int{1,24,5,6}
	// nums := []int{1,24,5,6}
	boolval := []bool{true, false, true}
	// names := []string{"golang", "typescript"}
	printSlice(boolval, "John")
	// printStringSlice(names)
}