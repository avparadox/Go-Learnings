package main

import "fmt"

func main() {
	// age := 16
	// if age >= 18 {
	// 	fmt.Println("Person is an adult.")
	// } else{
	// 	fmt.Println(("Person is not an adult"))
	// }

	marks := 98
	if marks >= 90 {
		fmt.Println("A+ Grade")
	} else if marks >= 75 && marks < 90 {
		fmt.Println("First Class")
	} else if marks >= 65 &&  marks <= 74 {
		fmt.Println("B Grade")
	} else if marks >= 36 && marks < 74 {
		fmt.Println("C Grade")
	} else{
		fmt.Println("Fail!")
	}
}