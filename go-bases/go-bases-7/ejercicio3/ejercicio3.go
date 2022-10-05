package main

import (
	"fmt"
	"os"
)

func main() {
	var salary int
	fmt.Println("Ingresa el salario a validar: ")
	fmt.Scan(&salary)
	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
