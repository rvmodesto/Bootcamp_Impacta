package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	fmt.Println(variavel1)

	variavel2 := "variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = 
		"opa"
		variavel4 string = "beleza?!"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}