package main

import "fmt"

func main() {
	fmt.Println("Iniciando o programa...")

	obterDivisao(2, 0)
}

func exibirRecover() {
	// Crio uma variavel recover e faço com que ela retome o programa
	if retomar := recover(); retomar != nil {
		fmt.Println("Programa retomado: ", retomar)
	}
}

func obterDivisao(numerador, denominador int) float64 {
	defer exibirRecover()
	if denominador == 0 {
		// O programa trava a partir do momento que ele enxerga o panic.
		panic("Não existe divisão por zero")
	}

	divisao := numerador / denominador

	return float64(divisao)
}
