package main

import (
	"fmt"
)

func main(){

	nome := func(palavra string) string{
		fmt.Sprintf("what's your name? %s", palavra)
		return palavra
	}("My name is Everson. Nice too meet you")

	fmt.Println(nome)


	
}