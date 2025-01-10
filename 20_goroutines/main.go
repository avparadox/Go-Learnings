package main

import (
	"fmt"
	"time"
)

// func task (id int){
// 	fmt.Println("Doing Task",id)
// }

func main(){
	for i := 0; i <= 20; i++{
	
		go func(i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}