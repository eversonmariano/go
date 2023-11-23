package main

import(
	"fmt"
)

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}


func alunoAprovado(n1, n2 float32) bool{
	
	defer recuperarExecucao()

	media := (n1+n2)/2
	
	if media > 7{
		return true
	} else if media < 7{
		return false
	}

	panic("A media é exatamente 7!")
}

func main(){
	fmt.Println(alunoAprovado(6,8))
	fmt.Println("pos execução")

}