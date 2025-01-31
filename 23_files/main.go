package main

import (
	"fmt"
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

	// read file - 1st Method
	// f, err := os.Open("example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// buff := make([]byte, 110)
	// d, err := f.Read(buff)
	// if err != nil{
	// 	panic(err)
	// }

	// for i := 0 ; i < len(buff) ; i++{
	// 	println("data", d,string(buff[i]))
	// }

	// read file - 2nd Method
	// should use only for small size of files
	// data, err := os.ReadFile("example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println(string(data))


	// reading folders

	// dir, err := os.Open("../")
	// if err != nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)

	// for _, fi := range fileInfo{
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// writing in files
	f, err := os.Create("example2.txt")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	// f.WriteString("hi golang!")

	// bytes := []byte("Hello Go")
	// f.Write(bytes)

	// read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil{
	// 		if err.Error() != "EOF"{
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if err != nil{
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written to new file successfully!")

	// Deleting a file

	os.Remove("example2.txt")

	fmt.Println("file removed")
}