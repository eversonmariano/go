package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("arrays e slices")

	var array1 [5]string
	array1[0] = "1"
	fmt.Println(array1)
	
	array2 := [5]string{"1","2","3","4","5"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(array3)

	slice := []int{10,11,12,13,14,15,16,17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[2:5]
	fmt.Println(slice2)

	array2[3] = "posiçao alterada"
	fmt.Println(slice2)

//ARRAY INTERNO

	fmt.Println("************************")
	slice3 := make([]float32, 10,11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) //capacity




	


}