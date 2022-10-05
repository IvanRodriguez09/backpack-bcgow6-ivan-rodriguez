package main

import (
	"fmt"
	"os"
)

type salaryError struct {
	msg string
}

func (e salaryError) Error() string {
	return e.msg
}

func main() {
	var salary int
	fmt.Println("Ingresa el salario a validar: ")
	fmt.Scan(&salary)
	if salary < 150000 {
		err := salaryError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
