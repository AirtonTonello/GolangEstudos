package main

import "fmt"

func main() {
	fmt.Println("Funções")

	name := getName()
	fmt.Println(name)

	name, age := getNameAndAge()
	fmt.Println(name, age)

	name, _, other := getNameAndAgeAndOther()
	fmt.Println(name, other)
}

func getName() string {
	name := "Airton"
	return name
}

func getNameAndAge() (string, int) {
	age := 28
	return getName(), age
}

func getNameAndAgeAndOther() (string, int, string) {
	other := "Outro"
	name, age := getNameAndAge()
	return name, age, other
}
