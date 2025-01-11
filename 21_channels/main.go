package main

import "fmt"

//sending
// func processNum (numChan chan int){
// 	for num := range numChan{
// 		fmt.Println("Processing", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

//receive
// func sum (result chan int, num1 int, num2 int){
// 	finalResult := num1 + num2
// 	result <- finalResult
// }

// goroutine synchronizer
func task (done chan bool){
	defer func ()  {
		done <- true	
	}()
	fmt.Println("Processing...")
}

func main(){

	done := make(chan bool)
	go task(done)

	<- done // blocking


	// result := make(chan int)
	// go sum(result,4,5)

	// res := <- result // blocking

	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)

	// for{
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)
	// messageChan <- "ping & pong" // Blocking Operation
	// msg := <-messageChan
	// fmt.Println(msg)
}