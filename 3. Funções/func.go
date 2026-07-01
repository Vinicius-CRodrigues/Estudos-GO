package main

import (
	"fmt"
	"math"
)

func somar(num1 int, num2 int) int {
	return num1 + num2

}

func dividir(num1, num2 float32) float32 {
	return num1 / num2
}

func calculoMatematico(num1, num2 int) (int, int, int) {

	somar := num1 + num2
	multiplicar := num1 * num2
	subtrair := num1 - num2

	return somar, multiplicar, subtrair

}

// RETORNO NOMEADO - Atribuimos um nome aos retornos das funções.
func calculoPotencia(num1, num2 int) (potencia int, raiz float64) {
	potencia = int(math.Pow(float64(num1), float64(num2)))
	raiz = math.Sqrt(float64(num1))

	return potencia, raiz
}

// FUNÇÕES VARIÁTICAS - Com os '...' criamos um slice que guarda varios valores em um função.
func escrevendo(nome string, idade ...uint) {
	for _, idade := range idade {
		fmt.Println(nome, idade)
	}
}

func main() {
	resultadoSoma := somar(3, 2)
	resultadoDivisao := dividir(3, 2)
	fmt.Println(resultadoSoma)

	// Com o underline, o terceiro elemento não precisa ser declarado.
	resultadoNovaSoma, resultadoMultiplicacao, _ := calculoMatematico(10, 20)

	fmt.Println("Soma:", resultadoNovaSoma)
	fmt.Println("Multiplicação:", resultadoMultiplicacao)
	fmt.Println(resultadoDivisao)

	//Criando funções dentro da função main.
	text := func(txt string) {
		fmt.Println(txt)
	}

	// Criando funções anonimas
	newText := func(texto string) string {
		return fmt.Sprintf("RETORNO -> %s", texto)

	}("Vinicius Cardoso")

	fmt.Println(newText)

	text("Testando a função dentro do main como um tipo")

	fmt.Println(calculoPotencia(3, 2))

	escrevendo("Indice", 1, 2, 3, 4, 5, 6)
}
