package main

import (
	"fmt"
)

type endeco struct {
	cidade string
	estado string
	bairro string
}

type usuario struct {
	nome           string
	idade          uint8
	altura         float32
	ativoNoSistema bool
	// Criando uma variável do tipo endereco que representa outra struct.
	endeco endeco
}

// Criando um mecnismo de herança aonde eu herdo as variaveis da struct #usuario
type estudante struct {
	usuario
	curso     string
	faculdade string
}

func main() {
	// Declaração explicita do struct
	var user1 usuario
	user1.nome = "Vini"
	user1.idade = 24
	user1.altura = 1.75

	fmt.Println(user1)

	// Declaração implicita do struct
	user2 := usuario{
		nome:           "Vini",
		idade:          24,
		altura:         1.75,
		ativoNoSistema: true,
		//Chamando a variável do tipo endereço.
		endeco: endeco{cidade: "Samambaia", estado: "Brasília/DF", bairro: "Samabaia sul"},
	}

	fmt.Println(user2)

	estudante1 := estudante{
		// Consigo pegar um usuário criado herdando dentro de uma nova struct.
		user2,
		"Engenharia de computação",
		"UnB",
	}
	fmt.Println(estudante1)

}
