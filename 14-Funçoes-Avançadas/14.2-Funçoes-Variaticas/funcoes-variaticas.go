package main

import (
	"fmt"
)

func soma( num ...int) int{
	total := 0
	for _, valor := range num{
		total += valor
	}
	return total
}

func escrever(texto string, num ...int){
	for _, val := range num{
		fmt.Println(texto,val)
	}
	return 
}


func main(){
	fmt.Println("total =", soma(1,2,3,4,10,30))
	
	escrever("oi mundo", 123,12,1)
}