package main

import (
	"fmt"
)

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multipĺicacao := 10 / 4
	divisao := 10 * 3
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multipĺicacao, divisao, restoDaDivisao)

	var num1 int16 = 10
	var num2 int16 = 2
	fmt.Println(num1 + num2)

	//OPERADORES DE ATRIBUICAO
	var var1 string = "oi"
	var2 := "tudo bem?"
	fmt.Println(var1, var2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNARIOS

	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero - 20

	numero *= 3 //numero = numero *3

	numero /= 5 //numero = numero / 5
	fmt.Println(numero)

}
