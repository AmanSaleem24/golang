package main

import (
	"fmt" 

	"github.com/AmanSaleem24/golang/auth"
	"github.com/AmanSaleem24/golang/user"
	"github.com/fatih/color"
)

func main(){
	auth.LoginWithCredentials("aman", "secretPassword")
	fmt.Println(auth.GetSession())

	user := user.User{
		Email: "aman@gmail.com",
		Name: "Aman Saleem",
	}

	
	fmt.Println(user)
	color.Red(user.Email)
	color.Green(user.Name)
}
 