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

	fmt.Println("file name: ", fileInfo.Name())
	fmt.Println("file true or dir false: ",fileInfo.IsDir())
	fmt.Println("file size: ",fileInfo.Size(), " bytes")
	fmt.Println("file permissions: ",fileInfo.Mode())
	fmt.Println("file modified: ",fileInfo.ModTime())


}