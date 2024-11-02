package main

// import "fmt"

// for is the only construct in Golang
func main(){
    // making a while loop
    // i:= 1
    // for i <= 3{
    //     fmt.Println(i)
    //     i++  // Don't forget this
    // }

    // Running an infite loop
    // for{
    //     // You can make a print statement without fmt as below.
    //     println("Aditya")
    //     println(1)
    // }

    // For loops
    // for i:= 0; i <= 3; i++ {
    //     if i == 2{
    //         continue
    //     }
    //     println(i)
    // }

    for i:= range 10{
        println("This number is", i)
    }
}