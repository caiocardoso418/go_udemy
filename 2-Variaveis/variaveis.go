package main

import "fmt"

func main() {
	//forma 1 de usar variaveis string
	var variavel1 string = "variavel 1"
	//forma 2 de usar variaveis string 
	variavel2 := "variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "lalalala"
		variavel4 string = "lalalala"
	)

	fmt.Println(variavel3,variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5,variavel6)
//trocando os dados das variaveis 
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5,variavel6)
}