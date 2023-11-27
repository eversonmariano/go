package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main(){
	generica(123)
	generica("true")
	generica(false)


	mapa := map[interface{}] interface{}{
		1: 123,
		float32(12.5): "japa",
		true: false,
		
	}
	fmt.Println(mapa)
}