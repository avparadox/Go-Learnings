package main

import (
	"fmt"
	"maps"
)

func main(){

	// Defining an Element
	m := make(map[string]string)
	// fmt.Println(m)
	
	// Setting an Element
	m["name"] = "Aditya"
	m["surname"] = "Vyas"
	// fmt.Println(m)

	// get an element
	// fmt.Println(m)
	// fmt.Println(m["name"], m["surname"])
	// fmt.Println(len(m))

	//Imp
	// fmt.Println(m["age"])
	// if key doesn't exists in the map, it returns zero value like above example.

	m2 := make(map[string]int)
	m2["age"] = 21
	m2["pincode"] = 777777 
	// String ->      , Int -> 0, Bool -> false
	// fmt.Println(m2["age"], m2["phone_number"])

	// Delete Function
	// delete(m2, "age")
	// fmt.Println(m2)

	clear(m2)
	// fmt.Println(m2)

	// One of the ways to create a map
	m3 := map[string]int{"price": 30}
	// fmt.Println(m3)

	// Maps Checking
	err, ok := m3["prie"]
	if ok {
		fmt.Println("Cool hai Ji")
	}else{
		fmt.Println(err)
	}
	// fmt.Println(key)

	// Maps Equality
	m4 := map[string]int{"price": 30}
	m5 := map[string]int{"price": 33}

	fmt.Println(maps.Equal(m4,m5))
}


