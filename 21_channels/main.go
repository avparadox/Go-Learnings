package main

import (
	"fmt"
	"time"
)

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
// func task (done chan bool){
// 	defer func ()  {
// 		done <- true
// 	}()
// 	fmt.Println("Processing...")
// }


//Email Channel
// func emailSender(emailChan chan string, done chan bool){

// 	defer func ()  {
// 		done <- true	
// 	}()

// 	for email := range emailChan{
// 		fmt.Println("Sending emails to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main(){

	// Email Channel
	// emailChan := make(chan string, 100)

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 6 ; i ++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("Done Processing")
	// close(emailChan)
	// <- done
	// emailChan <- "1@demo.com"
	// emailChan <- "2@demo.com"


	// done := make(chan bool)
	// go task(done)

	// <- done // blocking


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