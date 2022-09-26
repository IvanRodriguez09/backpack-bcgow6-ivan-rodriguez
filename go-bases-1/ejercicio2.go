package main

import "fmt"

func main() {
	var temperature float64 = 11.0
	var moisture int = 84
	var pressure float64 = 565
	fmt.Printf("La temperatura en bogotá es: %v°C\n", temperature)
	fmt.Printf("La humedad en bogotá es: %v%%\n", moisture)
	fmt.Printf("La presion atmosferica en bogotá es: %vmmHg\n", pressure)
}