package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var salary int
	fmt.Println("Ingresa el salario a validar: ")
	fmt.Scan(&salary)
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
