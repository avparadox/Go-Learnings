package main

import "fmt"

func add(a,b int) int{
	return a+b
}

func getLangauges()(string,string, int){
	return "golang", "JS", 34
}

// func processIt(fn func(a int)int){
// 	fn(1)
// }

func processIt() func(a int) int {
	return func (a int) int {
		return 4
	}
}


func main(){
	
	total := add(3,2)
	fmt.Println(total)

	// fmt.Println(getLangauges())
	l1, l2 , _ := getLangauges()
	fmt.Println(l1,l2)

	// fn := func (a int ) int {
	// 		return 2
	// }

	fn := processIt()
	fn(7)
}