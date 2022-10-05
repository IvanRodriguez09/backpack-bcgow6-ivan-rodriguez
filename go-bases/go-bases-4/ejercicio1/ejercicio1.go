package main

import (
	"fmt"
)

type student struct {
	name		string
	lastName	string
	dni			string
	date		string
}

func (s student) Detail() string{
	return fmt.Sprintf("My fullname is %s %s\nmy DNI is %s \nand my date of admission was %s \n",
		s.name, s.lastName, s.dni, s.date,
	) 
}

func main() {

	newStudent := student {
		name:		"Ivan",
		lastName: 	"Rodriguez",
		dni:		"123455",
		date:		"14/10/2022",
	}

	fmt.Println(newStudent.Detail())
}