package main

import (
	"fmt"
	"os"
)

func main(){
	f, err := os.Open("example.txt")
	if err != nil {
		// error handling / logging the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		// error handling / logging the error
		panic(err)
	}

	fmt.Println(fileInfo.Name())


}