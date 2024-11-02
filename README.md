# Coder's Gyan Golang Series

![Image](./image//YmGp5Uzh4ag-HD.jpg)

### Link to Playlist: [Happy Coding](https://www.youtube.com/playlist?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6)

## Chapter 1:

Installing Golang

```bash
sudo apt update
sudo apt install golang -y
```

## Chapter 2:

Hello World Application

Code Block:

```bash
package main

import "fmt"

func main(){
	fmt.Println("Hello World")
}
```

How to run the code:

`First get to the respective folder having go file, then run the below command`

```bash
go run main.go
```

## Chapter 3:

In this chapter we look up on simple values

Code Block:

```bash
package main

import (
	"fmt"
)

func main(){
	// Chapter 3

	// Simple Values

	// Integer
	fmt.Println(1+1)

	// Strings
	fmt.Println("Hello Golang")

	// Booleans
	fmt.Println(true)
	fmt.Println(false)

	//Floats
	fmt.Println(10.58)

	// Division
	fmt.Println(14.0/7.0)
}
```

How to run the code:

`First get to the respective folder having go file, then run the below command`

```bash
go run main.go
```

## Chapter 4:

In this chapter we focus on different varibales (int, bool, string), shorthand syntax & types of variables like int32, int 64 etc.

Code Block:

```bash
package main

import (
	"fmt"
)

func main(){
	// var name string = "Aditya"

	// Golang Infers the name type
	// var name = "Aditya"
	// fmt.Println("Hello",name)

	var is_adult = true;
	fmt.Println( is_adult)

	var age int = 23
	fmt.Println(age)

	// Shorthand Syntax
	name := "Aditya"
	fmt.Println(name)

	// Different Scenes
	var age2 int
	age2 = 244
	fmt.Println(age2)

	// Float Example
	var price float32
	price = 22.45
	fmt.Println(price)

	var price2 = 22.56
	fmt.Println(price2)

	price3 := 40.67
	fmt.Println(price3)
}
```

How to run the code:

`First get to the respective folder having go file, then run the below command`

```bash
go run variables.go
```

## Chapter 5:

In this chapter we look upon constants & it's way of writing.

Code Block:

```bash
package main

import "fmt"

const age = 23
const name = "aditya"


func main(){
	const name = "Aditya"
	// name = "Adi"  (Will give an err as constants cannot be changed)
	fmt.Println(name)
	fmt.Println(age)

	const (
		short_name = "adi"
		age = 21
	)
	fmt.Println(short_name)
}
```

How to run the code:

`First get to the respective folder having go file, then run the below command`

```bash
go run constants.go
```
## Chapter 6:

In this chapter we look upon For Loops & Range in Golang.

Code Block:

```bash
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

```