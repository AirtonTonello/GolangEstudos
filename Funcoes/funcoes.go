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

	name = setNameAndReturn(name)
	fmt.Println("Return:", name)

	test := Test{value: 10}
	compareResult := test.compare(20)
	fmt.Println("Compare", compareResult)

	test.value = 20
	compareResult = test.compare(20)
	fmt.Println("Compare", compareResult)

	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
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

func setNameAndReturn(name string) string {
	fmt.Println("Set:", name)
	return name
}

type Test struct {
	value int
}

/*
	receptor funciona como uma extensão no C#
	static int Compare(this int right, int left)

	caso o receptor não tenha o * significa que uma copia do objeto sera criada, igual o Ref no C#
*/
func (t *Test) compare(right int) int {
	if t.value == right {
		return 1
	}

	return 0
}

/*
	numero indeterminado de parametros
	igual o param no C#
*/
func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}
