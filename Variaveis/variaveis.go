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

	var array [4]string
	array[0] = "0"
	array[1] = "1"
	array[2] = "2"
	array[3] = "3"

	fmt.Println("Array Value:", array)
	fmt.Println("Type Of:", reflect.TypeOf(array))

	fmt.Println("---------------------------------------------------")

	list := []string{"0", "1", "2", "3"}
	list = append(list, "4")

	fmt.Println("Slice Value:", list)
	fmt.Println("Type Of:", reflect.TypeOf(list))

	fmt.Println("---------------------------------------------------")

	var slices []int

	for i := 0; i < 10; i++ {

		fmt.Println("Slice Value:", slices)
		fmt.Println("Type Of:", reflect.TypeOf(slices))
		fmt.Println("Length:", len(slices))
		fmt.Println("Capacity:", cap(slices))

		slices = append(slices, i)

		fmt.Println()
	}

	// o slice aumenta a capacidade de acordo com o tamanho do ultimo length
	// ele não aumenta a capacidade de acordo com o valor que foi iniciado

	fmt.Println("---------------------------------------------------")
}
