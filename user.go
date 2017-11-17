package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

func (u *User) registration(){
	fmt.Println("Registration")

	fmt.Print("Please tell me your name: ")
	fmt.Scanln(&u.Name)

	fmt.Print("Your e-mail: ")
	fmt.Scanln(&u.Email)

	fmt.Print("Your phone: ")
	fmt.Scanln(&u.Phone)

	fmt.Print("Your password: ")
	fmt.Scanln(&u.Password)
}

func (u User) print() {
	fmt.Println("name: ", u.Name)
	fmt.Println("email: ", u.Email)
	fmt.Println("phone: ", u.Phone)
	fmt.Println("password: ", u.Password)
}
