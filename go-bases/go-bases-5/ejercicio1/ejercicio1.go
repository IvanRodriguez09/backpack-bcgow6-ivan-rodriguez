package main

import (
	"fmt"
	"os"
)

type Product struct {
	id int
	price float64
	amount int
}

func (p Product) details() string{
	return fmt.Sprintf("%v,%v,%v\n", p.id, p.price, p.amount)
}

func main() {
	product := Product {id: 1, price: 1000, amount: 50}
	product2 := Product {id: 2, price: 400, amount: 20}
	product3 := Product {id: 3, price: 1500, amount: 5}
	content := []byte(fmt.Sprint(product.details(),product2.details(),product3.details()))
	fmt.Println(content)
	err := os.WriteFile("./myProducts.csv", content, 0644)
	if err != nil {
		fmt.Println("Something was wrong")
	} else {
		fmt.Println("File wrote correctly")
	}
	

}