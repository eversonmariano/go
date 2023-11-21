package main

import (
	"fmt"
)

func main(){
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1,var2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var var3 int 
	var var4 *int

	var3 = 20
	var4 = &var3 //guarda o endereço de memória de var3
	var3++

	
	fmt.Println(var3, *var4) //*var4 pega o valor contido no endereço de memória de var3
	
	
}
