package main

import "fmt"


type usuario struct {
	nome string
	idade uint8
}

func (u *usuario) fazerAniversario(){
	u.idade++

}

func (u usuario) salvar(){
	fmt.Println("Salvando os dados do Usuário %s no banco de dados.", u.nome)
	fmt.Println("Salvando os dados do Usuário %d no banco de dados.", u.idade)
}

func main(){
	usuario1 := usuario{"everson", 40}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"beta", 45}
	fmt.Println(usuario2.idade)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
	

}