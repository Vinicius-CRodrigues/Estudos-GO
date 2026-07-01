package main

import "fmt"

func main() {
	verdadeiro, falso, numero := true, false, 1

	// Operador lógico OR
	fmt.Println("Resultado 1: ", verdadeiro || falso)

	// Operador lógico NOT
	fmt.Println("Resultado 2: ", !verdadeiro)

	// Operador lógico AND
	fmt.Println("Resultado 3: ", verdadeiro && falso)

	numero++
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero += 1
	fmt.Println(numero)

	numero *= 2
	fmt.Println(numero)

	numero /= 1
	fmt.Println(numero)

	numero %= 1
	fmt.Println(numero)
}
