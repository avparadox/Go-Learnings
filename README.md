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
