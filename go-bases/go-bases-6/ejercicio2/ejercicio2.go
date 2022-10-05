package main

import (
	"fmt"
)

type Product struct {
	name	string
	price	float64
	amount	int
}

type User struct {
	name		string
	lastName	string
	email		string
	products	[]Product
}

func newProduct(name string, price float64) (product Product){
	product = Product {
		name: 	name,
		price:	price,
	}
	fmt.Println("Product created successfully")
	return 
}

func newUser(name string, lastName string, email string) (user User){
	user = User {
		name: 		name,
		lastName: 	lastName,
		email:		email,
	}
	fmt.Println("User created successfully")
	return 
}

func addProduct(user *User, product *Product, amount int) {
	product.amount = amount
	user.products = append(user.products, *product)
	fmt.Println("Product added successfully")
}

func removeProducts(user *User) {
	user.products = user.products[:0]
	fmt.Printf("%+v\n", user)
	fmt.Println("Product removed successfully")
}

func main() {
	
	user := newUser("Ivan", "Rodriguez", "ivan@email.com")
	product := newProduct("Juice", 1500)
	product2 := newProduct("All Rich", 2500)
	addProduct(&user, &product, 2)
	addProduct(&user, &product2, 5)
	fmt.Printf("%+v\n", user)
	removeProducts(&user)
}