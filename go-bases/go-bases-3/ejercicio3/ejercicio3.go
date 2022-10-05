package main

import "fmt"

const (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
)

func getSalary(minutes float64, category string) (salary float64) {
	
	hours := minutes/60

	switch category {
	case CategoryA:
		salary = hours * 3000
		salary *= 1.5
	case CategoryB:	
		salary = hours * 1500
		salary *= 1.2
	case CategoryC:
		salary = hours * 1000
	}
	return
}

func main()  {

	fmt.Println(getSalary(2400, "A"))
	fmt.Println(getSalary(2400, "B"))
	fmt.Println(getSalary(2400, "C"))
	
}