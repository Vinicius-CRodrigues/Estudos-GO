package main

import "fmt"

func main() {

	// Aqui criamos
	var variavel int
	var pointer *int

	// Aqui eu consegui capturar o endereço de memória da variável numUm.
	pointer = &variavel

	fmt.Printf("Valor da variável: %d, endereço em que ela se encontra: %p \n", variavel, pointer)

	variavel += 10

	ValorSalvoNoPointer := *pointer // desfazendo a referencia.

	fmt.Printf("Valor salvo no pointer: %d, valor da variável: %d \n", ValorSalvoNoPointer, variavel)

}
