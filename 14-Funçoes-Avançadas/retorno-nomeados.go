package main

import (
	"fmt"

	)

func calMatematicos(n1, n2 int)(soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main(){
	soma, subtracao := calMatematicos(20, 5)
	fmt.Println(soma, subtracao)
}