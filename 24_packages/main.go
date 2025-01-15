package main

import (
	"fmt"

	"github.com/avparadox/Go-Learnings/auth"
	"github.com/avparadox/Go-Learnings/user"
)

func main(){
	auth.LoginWithCredentials("Aditya", "secret")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "user@email.com",
		Name: "John Doe",
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

}