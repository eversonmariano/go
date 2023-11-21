package main

import(
	"fmt"
)

func diaDaSemana(num int) string {
	switch num{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Num invalido"
	
	}

}

func diaDaSemana2(num int)string{
	var diaDaSemana string
	switch {
		case num == 1:
			diaDaSemana = "Domingo"
		case num == 2:
			diaDaSemana = "Segunda"
		case num == 3:
			diaDaSemana = "Terça"
		case num == 4:
			diaDaSemana = "Quarta"
		case num == 5:
			diaDaSemana = "Quinta"
		case num == 6:
			diaDaSemana = "Sexta"
		case num == 7:
			diaDaSemana = "Sabado" 
		default:
			diaDaSemana = "Inválido"
				
	}
	return diaDaSemana
} 

func main(){
	
	n := 7 
	if n >= 1 || n <= 7{
		fmt.Println(diaDaSemana(n))
	}else{
		fmt.Println("Numero inválido.")
	} 

	fmt.Println("--------------------")

	x := 7
	fmt.Println(diaDaSemana2(x))
	


}