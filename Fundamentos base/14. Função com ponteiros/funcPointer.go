package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1

}

func inverteSinalComPointer(numero *int) {
	// Pega o conteúdo de numero e multiplica por menos 1
	*numero = *numero * -1

}

func arrPointer(arr *[5]int) {
	*arr = [5]int{1, 2, 4, 5, 6}
}

func main() {
	numero := 10

	// Faço somente uma cópia de numero, criando uma nova variável.
	inverte := inverteSinal(numero)

	fmt.Println("Invertido sem ponteiro:", inverte)
	fmt.Println("Numero da variável inalterado:", numero)

	novoNumero := 13

	// Captura o endereço de memória de uma variável já existente.
	inverteSinalComPointer(&novoNumero)
	fmt.Println("Invertido com ponteiro:", novoNumero)

	var arr [5]int

	arrPointer(&arr)

	fmt.Println("Array preenchido com ponteiro:", arr)
}

// Uma função que é inicializada antes da main
func init() {
	fmt.Println("Inicializado o programa...")
}
