package main

import (
	"os"
)

func main(){
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// error handling / logging the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// error handling / logging the error
	// 	panic(err)
	// }

	// fmt.Println("file name: ", fileInfo.Name())
	// fmt.Println("file true or dir false: ",fileInfo.IsDir())
	// fmt.Println("file size: ",fileInfo.Size(), " bytes")
	// fmt.Println("file permissions: ",fileInfo.Mode())
	// fmt.Println("file modified: ",fileInfo.ModTime())
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// error handling / logging the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// error handling / logging the error
	// 	panic(err)
	// }

	// fmt.Println("file name: ", fileInfo.Name())
	// fmt.Println("file true or dir false: ",fileInfo.IsDir())
	// fmt.Println("file size: ",fileInfo.Size(), " bytes")
	// fmt.Println("file permissions: ",fileInfo.Mode())
	// fmt.Println("file modified: ",fileInfo.ModTime())

	// read file
	f, err := os.Open("example.txt")
	if err != nil{
		panic(err)
	}

	defer f.Close()

	buff := make([]byte, 110)
	d, err := f.Read(buff)
	if err != nil{
		panic(err)
	}

	for i := 0 ; i < len(buff) ; i++{
		println("data", d,string(buff[i]))
	}

}