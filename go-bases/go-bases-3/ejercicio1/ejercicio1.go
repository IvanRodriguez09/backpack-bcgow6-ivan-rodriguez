package main

import "fmt"

func getSalaryTax(salary float64) float64{
	if salary > 150000 {
		return salary*0.27
	}else if (salary <= 150000 && salary > 50000) {
		return salary*0.17
	} 
	return 0
}

func main()  {
	fmt.Println(getSalaryTax(100000))
	fmt.Println(getSalaryTax(1000000))
	fmt.Println(getSalaryTax(17000))
	fmt.Println(getSalaryTax(50000))
}