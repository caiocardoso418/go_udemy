package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 100
	fmt.Println(numero)

	var numeroreal float32 = 123.45 //tendo o 64 
	fmt.Println(numeroreal)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B' //tabela ascc
	fmt.Println(char)

	var texto int16 //valor 0 
	fmt.Println(texto)

	var bolleano1 bool = false
	fmt.Println(bolleano1)

	var erro error = errors.New("Erro interno") //é um tipo erro e nao string e tem que importar o erros
	fmt.Println(erro)
}
