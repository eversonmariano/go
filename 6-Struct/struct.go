package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
	endereco endereco 
}

type endereco struct {
	rua string
	num uint16
}

func main(){
	fmt.Println("ARQUIVO STRUCT")

	var u usuario //chamei a struct para dentro da main
	u.nome = "Everson"
	u.idade = 40

	fmt.Println(u)

	endereco1 := endereco{"Rua das Flores", 525}


	usuario2 := usuario{"Beta", 45, endereco1}
	fmt.Println(usuario2)


}