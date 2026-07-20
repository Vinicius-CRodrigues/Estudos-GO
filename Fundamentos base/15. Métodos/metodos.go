package main

import "fmt"

type usuario struct {
	nome   string
	idade  int
	altura float32
	peso   float32
}

// Criando um método para fazer a manipulação do struct
func (user usuario) exibirDados() {
	fmt.Println("Nome:", user.nome)
	fmt.Println("Idade:", user.idade)
	fmt.Println("Altura:", user.altura)
	fmt.Println("peso:", user.peso)

}

func (user usuario) calcularIMC() {
	IMC := user.peso / (user.altura * user.altura)

	if IMC < 18.5 {
		fmt.Printf("%s ESTA ABAIXO DO PESO NORMAL !\n", user.nome)
	} else if IMC >= 18.5 && IMC <= 24.9 {
		fmt.Printf("%s ESTA COM PESO NORMAL !\n", user.nome)
	} else if IMC >= 25 && IMC <= 29.9 {
		fmt.Printf("%s ESTA COM EXCESSO DE PESO !\n", user.nome)
	} else if IMC >= 30 && IMC <= 34.9 {
		fmt.Printf("%s ESTA COM OBESIDADE CLASSE 1 !\n", user.nome)
	} else if IMC >= 30 && IMC <= 39.9 {
		fmt.Printf("%s ESTA COM OBESIDADE CLASSE 2 !\n", user.nome)
	} else {
		fmt.Printf("%s ESTA COM OBESIDADE CLASSE 3 !\n", user.nome)
	}
}

// Usando métodos com ponteiros para alterar as variáveis do struct
func (user *usuario) fazerAniversario() {
	fmt.Printf("%s FEZ ANIVERSÁRIO !\n", user.nome)
	user.idade += 1
}

func main() {
	usuario1 := usuario{"Vinicius", 24, 1.74, 80}

	usuario1.exibirDados()
	usuario1.calcularIMC()

	usuario1.fazerAniversario()
	usuario1.exibirDados()
}
