package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8){ //pode retornar dois bagui
	soma := n1+n2
	subtaracao := n1 -n2
	return soma, subtaracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da funcao 1")
	fmt.Println(resultado)

	resultadosSoma, _ := calculosMatematicos(10,15) //o _ pode tirar a variavel de la sem reclamar
	fmt.Println(resultadosSoma)
}