package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := -10000000
	fmt.Println(numero)

	var num2 uint32 = 10
	fmt.Println(num2)

	// alias do uint(inteiro sem sinal negativo)
	// int32 = rune
	//byte = int8

	var num3 rune = 88888
	fmt.Println(num3)

	var num4 byte = 2
	fmt.Println(num4)

	//NUMEROS REAIS
	var numReal1 float32 = 123.45
	fmt.Println(numReal1)

	var numReal2 float64 = 123.45
	fmt.Println(numReal2)

	numReal3 := 123.45
	fmt.Println(numReal3)

	//booleanos
	var boolean1 bool = true
	fmt.Println(boolean1)

	//tipo error

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

	
}
