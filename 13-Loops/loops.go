package main

import (
	"fmt"
)

func main() {
	// fmt.Println("loops")
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("implementando i=",i)
	// }

	// for j := 0; j < 10; j++{
	// 	fmt.Println("implementando j=", j)
	// }

	// nomes := [3]string{"ell", "snoopy", "beta"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	usuario := map[string]string{
		"nome":      "ell",
		"sobrenome": "mariano",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("executando infinitamente")
	}

}
