package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Product struct {
	id int
	price float64
	amount int
}

func getProduct(line string) (product Product) {
	
	details := strings.Split(line, ",")
	idValue, err := strconv.Atoi(details[0])
	priceValue, err2 := strconv.ParseFloat(details[1], 8)
	amountValue, err3 := strconv.Atoi(details[2])
	if (err != nil || err2 != nil || err3 != nil ) {
		fmt.Println("Something wrong with the registry")
	} else  {
		product = Product{
			id: idValue,
			price: priceValue,
			amount: amountValue,
		}
	}
	return product
}

func registryProduct(product Product, header *bool) {
	if !*header {
		line := fmt.Sprintf("%v\t\t%.2f\t\t%d\n",product.id, product.price, product.amount)
		fmt.Printf(line)
	} else {
		line := fmt.Sprintf("ID\t\tPRICE\t\tAMOUNT\n")
		fmt.Printf(line)
		line = fmt.Sprintf("%v\t\t%.2f\t\t%d\n",product.id, product.price, product.amount)
		fmt.Printf(line)
		*header = false
	}
}

func main() {
	data, err := os.ReadFile("./myProducts.csv")
	var total float64
	header := true
	if err != nil {
		fmt.Println("Something was wrong reading the file")
	} else {
		fdata := strings.Split(string(data), "\n")
		for _, line := range fdata {
			if(len(line) > 0){
				prod := getProduct(line)
				total += prod.price
				registryProduct(prod, &header)
			} else {
				fmt.Printf("\t\t%v\n", total)
			}
			
		}
	}
}