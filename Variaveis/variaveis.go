package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Variaveis")
	fmt.Println("---------------------------------------------------")

	var name string
	fmt.Println("String Default Init:", name)

	name = "Airton"
	fmt.Println("String Value:", name)

	fmt.Println("---------------------------------------------------")

	var version float32
	fmt.Println("Float Default Init:", version)

	version = 1.0
	fmt.Println("Float Value:", version)

	fmt.Println("---------------------------------------------------")

	var age int
	fmt.Println("Int Default Init:", age)

	age = 20
	fmt.Println("Int Value:", age)

	fmt.Println("---------------------------------------------------")

	var condition bool
	fmt.Println("Bool Default Init:", condition)

	condition = true
	fmt.Println("Bool Value:", condition)

	fmt.Println("---------------------------------------------------")

	var varString = "String"
	fmt.Println("Var Value:", varString)

	fmt.Println("---------------------------------------------------")

	var valueTypeOf = 10.00
	fmt.Println("Var Value:", valueTypeOf)
	fmt.Println("Type Of:", reflect.TypeOf(valueTypeOf))

	fmt.Println("---------------------------------------------------")

	// igual delphi
	value := "Short Var"
	fmt.Println("Infer Value:", value)
	fmt.Println("Infer Type:", reflect.TypeOf(value))

	fmt.Println("---------------------------------------------------")
}
