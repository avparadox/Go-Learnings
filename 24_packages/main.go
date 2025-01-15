package main

import (
	"fmt"

	"github.com/avparadox/Go-Learnings/auth"
	"github.com/avparadox/Go-Learnings/user"
	"github.com/fatih/color"
)

func main(){
	auth.LoginWithCredentials("Aditya", "secret")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@email.com",
		Name: "John Doe",
	}

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)
	color.Red(user.Email)

// go mod tidy - command to make mod file proper and install the remaining external packages 	

}