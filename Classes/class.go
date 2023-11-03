package main

import (
	"fmt"
	"time"
)

type Product struct {
	id          int
	description string
	price       float64
}

type Person struct {
	id          int
	firstName   string
	lastName    string
	dateOfBirth time.Time
}

func main() {
	notebook := Product{id: 1, description: "Notebook", price: 1500.00}
	fmt.Println(notebook)

	me := Person{1, "Airton", "Tonello", time.Now()}
	fmt.Println(me)

	var meAsPerson *Person
	meAsPerson = new(Person)
	meAsPerson.id = me.id + 1
	meAsPerson.firstName = me.firstName
	meAsPerson.lastName = me.lastName
	meAsPerson.dateOfBirth = me.dateOfBirth

	fmt.Println(meAsPerson)
	fmt.Println(*meAsPerson)

	var celular *Product
	celular = new(Product)
	celular.id = 2
	celular.description = "Motorola"
	celular.price = 1000.00

	fmt.Println(celular)
	fmt.Println(*celular)

	// comparações
	var carro1 Product = Product{1, "Carro", 10.00}
	var carro2 Product = Product{1, "Carro", 10.00}

	fmt.Println(carro1)
	fmt.Println(carro2)
	fmt.Println(carro1 == carro2)

	var carro3 *Product
	carro3 = new(Product)
	carro3.id = 2
	carro3.description = "Carro"
	carro3.price = 20.00

	var carro4 *Product
	carro4 = new(Product)
	carro4.id = 2
	carro4.description = "Carro"
	carro4.price = 20.00

	fmt.Println(*carro3)
	fmt.Println(carro3)
	fmt.Println(&carro3)
	fmt.Println(*carro4)
	fmt.Println(carro4)
	fmt.Println(&carro4)
	// compara posições na memoria
	fmt.Println(carro3 == carro4)
	// compara o conteudo da classe
	fmt.Println(*carro3 == *carro4)

	brinquedo := Product{}
	brinquedo.id = 1
	brinquedo.description = "Brinquedo"
	brinquedo.price = 55.00

	brinquedo.addTax(15.0)

	fmt.Println(brinquedo.price)
}

func (p *Product) addTax(tax float64) {
	p.price += tax
}
