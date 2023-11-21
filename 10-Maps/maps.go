package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"nome":      "Everson",
		"sobrenome": "Mariano",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiroNome": "Beta",
			"ultimoNome":   "Mariano",
		},
		"curso": {
			"nomeCurso":    "Enfermagem",
			"Universidade": "UEPB",
		},
	}
	fmt.Println(usuario2)
	

}
