package main

import "fmt"

func sum(n1, n2 int) int {
	return n1 + n2
}

func sub(n1, n2 int) int {
	return n1 - n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

func main() {
	fmt.Println("-- SWITCH CASE --")

	var (
		operacao string
		numUm    int
		numDois  int
	)

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&numUm)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&numDois)

	fmt.Print("Digite a operação: ")
	fmt.Scan(&operacao)

	switch operacao {
	case "+":
		fmt.Println("Resultado:", sum(numUm, numDois))

	case "-":
		fmt.Println("Resultado:", sub(numUm, numDois))

	case "*":
		fmt.Println("Resultado:", mult(numUm, numDois))

	default:
		fmt.Println("Resultado: Operação não encontrada")
	}

}
