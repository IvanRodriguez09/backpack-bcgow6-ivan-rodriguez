package main

import (
	"fmt"
	"time"
)

type Product struct {
	name   string
	price  float64
	amount int
}

type Service struct {
	name          string
	price         float64
	minutesWorked int
}

type Maintenance struct {
	name  string
	price float64
}

func productTotal(products []Product, c chan float64) {
	var total float64
	for _, product := range products {
		total += product.price * float64(product.amount)
	}
	c <- total

}

func serviceTotal(services []Service, c chan float64) {
	var total float64
	for _, service := range services {
		if service.minutesWorked%30 > 0 {
			total += service.price * float64((service.minutesWorked/30)+1)
		} else {
			total += service.price * float64((service.minutesWorked / 30))
		}
	}
	c <- total
}

func maintenanceTotal(maintenance []Maintenance, c chan float64) {
	var total float64
	for _, maint := range maintenance {
		total += maint.price
	}
	c <- total
}

func createElements() ([]Product, []Service, []Maintenance) {
	products := []Product{
		Product{
			name:   "Modem",
			price:  2000,
			amount: 10,
		},
		Product{
			name:   "Phone",
			price:  500,
			amount: 4,
		},
		Product{
			name:   "CPU",
			price:  1500,
			amount: 4,
		},
	}
	services := []Service{
		Service{
			name:          "Installation",
			price:         200,
			minutesWorked: 60,
		},
		Service{
			name:          "Cleaning",
			price:         500,
			minutesWorked: 70,
		},
	}
	maintenance := []Maintenance{
		Maintenance{
			name:  "Display",
			price: 1000,
		},
		Maintenance{
			name:  "Preventive",
			price: 2000,
		},
	}
	return products, services, maintenance
}
func main() {
	products, services, maintenance := createElements()
	//cProduct, cService, cMaintenance := make(chan float64), make(chan float64), make(chan float64)
	canal := make(chan float64)
	start := time.Now()
	var total float64
	//go productTotal(products, cProduct)
	//go serviceTotal(services, cService)
	//go maintenanceTotal(maintenance, cMaintenance)
	go productTotal(products, canal)
	go serviceTotal(services, canal)
	go maintenanceTotal(maintenance, canal)
	//total += <-cProduct
	//total += <-cMaintenance
	//total += <-cService
	total += <-canal
	total += <-canal
	total += <-canal
	finish := time.Now()
	fmt.Printf("Time total: %v:\tTotal is: %v\n", finish.Sub(start), total)
}
