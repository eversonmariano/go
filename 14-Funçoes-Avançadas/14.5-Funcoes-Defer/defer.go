package main

import(
	"fmt"
)

func funcao1(){
	fmt.Println("Entrando na funcao1")
}

func funcao2(){
	fmt.Println("Entrando na funcao2")
}

func alunoAprovado(n1,n2 float32) bool{
	defer fmt.Println("Media calculada. O resultado é..")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado...")
	
	media := (n1+n2)/2

	if media >= 7 {
		fmt.Println("Aluno aprovado")
		return true
	}
	return false
}

func main(){
	fmt.Println(alunoAprovado(7.5, 9))
}