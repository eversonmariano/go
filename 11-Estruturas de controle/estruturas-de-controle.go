package main

import(
	"fmt"
)

func main(){
fmt.Println("Estruturas de controle")

	num := 10

	if num > 15 {
		fmt.Println("num maior que 15")
	} else{
		fmt.Println("num menor ou igual a 15")
	}

	if outroNume := num; outroNume > 0{ // if_init -> cria uma variavel apenas para essa estrutura/condiçao e faz a verificação 
		fmt.Println("num menor ou igual a zero")
	} else if num < 10{
		fmt.Println("num menor ou igual a -10")

	} else{
		fmt.Println("entre -10 e 15")
	}





}