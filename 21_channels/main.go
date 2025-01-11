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

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func ()  {
		chan1 <- 10	
	}()

	go func() {
		chan2 <- "golang"
	}()

		for i := 0; i < 2; i++{
			select{
			case chan1Val := <- chan1:
				fmt.Println("received data from chan1", chan1Val)
			case chan2Val := <- chan2:
				fmt.Println("received data from chan2", chan2Val)
			}
		}




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