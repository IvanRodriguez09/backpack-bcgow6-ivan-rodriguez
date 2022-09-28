package main

import (
	"fmt"
)

type User struct {
	name		string
	lastName	string
	age			int
	email		string
	password	string
}

func updateName(user *User, newName string) {
	user.name = newName
	fmt.Println("the user name was updated successfully")
}

func updateLastName(user *User, newLastName string) {
	user.lastName = newLastName
	fmt.Println("the user last name was updated successfully")
}

func updateAge(user *User, newAge int) {
	user.age = newAge
	fmt.Println("the user age was updated successfully")
}

func updateEmail(user *User, newEmail string) {
	user.email = newEmail
	fmt.Println("the user email was updated successfully")
}

func updatePassword(user *User, newPassword string) {
	user.password = newPassword
	fmt.Println("the user password was updated successfully")
}

func (u User) details() {
	fmt.Printf("Name: %s\tLast name: %s\tage: %d\temail: %s\n",u.name, u.lastName, u.age, u.email)
}

func main()  {
	user := User {
		name: 		"Ivan",
		lastName:	"Rodriguez",
		age:		15,
		email:		"ivan@email.com",
		password: 	"12345",
	}
	user.details()
	updateName(&user, "Arturo")
	updateLastName(&user, "Pineda")
	updateAge(&user, 16)
	updateEmail(&user, "arturo@email.com")
	user.details()
}