package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println(arr)

	arr2 := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(arr2)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)

	fmt.Println(slice)

	// Diferenciando o tipo de dado:
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(arr))

	// Lembrando que o indice inicial e inclusivo e o final é exclusivo.
	novoSlice := arr[0:3]

	novoSlice[2] = 187

	fmt.Println(novoSlice)

	// arrays internos:
	sliceArrayInterno := make([]float32, 3, 4)
	sliceArrayInterno[0] = 10
	sliceArrayInterno[1] = 20
	sliceArrayInterno[2] = 30

	for i := 0; i < 10; i++ {
		sliceArrayInterno = append(sliceArrayInterno, float32(i))
	}

	fmt.Printf("Elementos: %v, tamanho: %d, capacidade: %d \n", sliceArrayInterno, len(sliceArrayInterno), cap(sliceArrayInterno))

}
