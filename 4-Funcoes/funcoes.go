package main

import(
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// funçao com dois retornos
func calculosMatematicos(n1, n2 int16) (int16, int16){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}



func main(){
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	result := f("texto da funcao f")
	fmt.Println(result)

	fmt.Println(calculosMatematicos(35, 10))

	resultSoma, resultSub := calculosMatematicos(10, 15)
	fmt.Println(resultSoma, resultSub)

	resultSoma1, _ := calculosMatematicos(10, 15)
	fmt.Println("so a soma = ", resultSoma1)

	_ , resultSub1 := calculosMatematicos(10, 15)
	fmt.Println("so a subtraçao = ",resultSub1)

	fmt.Println(calculosMatematicos(35, 15))
}