package main

import (
	"fmt"
	"os"
)

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return file
}

func main() {
	defer fmt.Println("ejecución finalizada")
	file := readFile("./customers.txt")
	fmt.Printf("%s\n", file)
}
