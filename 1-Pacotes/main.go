package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

//erro valida o email
func main() {
	fmt.Println("ESCREVENDO DO ARQUIVO MAIN.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("eversonmariano@yahoo.com.br")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("123")
	fmt.Println(erro2)

}
