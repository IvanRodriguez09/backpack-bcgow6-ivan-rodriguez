package main

import (
	"fmt"
	"errors"
) 

const (
	Dog = "dog"
	Cat = "cat"
	Hamster = "hamster"
	Tarantula = "tarantula"
)

func animalDog(amount int) float64 {
	return float64(amount)*10
}

func animalCat(amount int) float64 {
	return float64(amount)*5
}

func animalHamster(amount int) float64 {
	return float64(amount)*0.25
}

func animalTarantula(amount int) float64 {
	return float64(amount)*0.15
}

func Animal(name string) (func(amount int) float64, error)  {
	switch name {
	case Dog:
		return animalDog, nil
	case Cat:
		return animalCat, nil
	case Hamster:
		return animalHamster, nil
	case Tarantula:
		return animalTarantula, nil
	}
	return nil, errors.New("The animal isn't in the list")
}

func Animalorchestrator(animalType string, amount int) float64{
	animalOperator, err := Animal(animalType)
	if err != nil {
		fmt.Println(err)
	} else {
		return animalOperator(amount)
	}
	return 0
}

func main()  {
	var amount float64
	amount += Animalorchestrator(Dog, 8)
	amount += Animalorchestrator(Cat, 5)
	amount += Animalorchestrator(Hamster, 4)
	amount += Animalorchestrator(Tarantula, 10)
	fmt.Printf("The food necessary is: %v Kg \n",amount)
}