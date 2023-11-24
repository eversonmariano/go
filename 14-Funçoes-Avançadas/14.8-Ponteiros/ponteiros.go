package main

import(
	"fmt"
)

func inverterSinal(num int) int {
	return num * -1
}

func invertSinalComPonteiro(num *int){
	*num = *num * -1
}

func main(){
	numero := 20
	numSinalInvert := inverterSinal(numero)
	fmt.Println(numSinalInvert)
	println(numero)

	novoNum := 40
	fmt.Println(novoNum)
	invertSinalComPonteiro(&novoNum)
	fmt.Println(novoNum)
	


}