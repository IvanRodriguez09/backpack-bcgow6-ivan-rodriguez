package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Client struct {
	legajo      string
	name        string
	lastName    string
	dni         string
	phoneNumber string
	address     string
}

const (
	filename = "customers.txt"
)

func generateLegajo() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic("error generating legajo")
	}
	return string(newUUID)
}

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
	return file
}

func createClient(legajo string) (client Client) {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
		}
	}()
	client = Client{
		legajo:      legajo,
		name:        "Ivan",
		lastName:    "Rodriguez",
		dni:         "123456",
		phoneNumber: "32112345",
		address:     "the world",
	}
	fmt.Printf("%+v\n", client)
	return
}

func validateClientData(client *Client) bool {
	if client.name == "Ivan" {
		panic("Alguno de los valores ingresados esta vacio")
	}
	return true
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
		}
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()
	legajo := generateLegajo()
	client := createClient(legajo)
	readFile(filename)
	validateClientData(&client)

}
