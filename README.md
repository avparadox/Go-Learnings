# Coder's Gyan Golang Series

![Image](./image//YmGp5Uzh4ag-HD.jpg)

### Link to Playlist: [Happy Coding](https://www.youtube.com/playlist?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6)


##  Chapter 1:

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

```First get into the respective folder having go file, then run the below command```

``` bash 
go run main.go
```

## Chapter 3:

Here we looked upon simple values & printed them on the page

Code Block:

``` bash
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

```First get into the respective folder having go file, then run the below command```

``` bash 
go run main.go
```
 
## Chapter 4:

Here we focus on Different varibales (int, bool, string), shorthand syntax  & types of variables like int32, int 64 etc

Code Block:

``` bash
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

```First get into the respective folder having go file, then run the below command```

``` bash 
go run variables.go
```