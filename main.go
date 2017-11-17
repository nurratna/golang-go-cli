package main

import "fmt"

func main() {
	fmt.Println("Go-CLI")

	user := User{
		Name: "",
		Email: "",
		Phone: "",
		Password: "",
	}
	
	user.registration()
	user.Save("data/users.json")
	user.print()



}

