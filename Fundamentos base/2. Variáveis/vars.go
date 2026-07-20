package main

import "fmt"

func main() {
	//Declarando uma variável de forma explicita
	var nome string = "Vinícius"

	//Declarando uma variável de forma implicita
	idade := 24

	var (
		sexo   string  = "Masculino"
		altura float64 = 1.75
	)

	dataDeNascimento, diaDeNascimento, mesDeNascimento := 2001, 22, 12

	// Podemos alterar os conteúdos do mesmo tipo.
	sexo, nome = nome, sexo

	fmt.Printf("Meu nome é %s e a minha idade é %d anos. Minha altura é de %.2f	 e meu sexo é %s\n", nome, idade, altura, sexo)
	fmt.Printf("Minha data de nascimento é %d/%d/%d\n", diaDeNascimento, mesDeNascimento, dataDeNascimento)

}
