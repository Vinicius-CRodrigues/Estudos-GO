package main

import (
	"fmt"
)

func leituraMapa(mapaFrutas map[string]float32) {
	for chave, valor := range mapaFrutas {
		fmt.Printf("Fruta: %s. Preço: R$ %.2f \n", chave, valor)
	}
}

func leituraMapaMapa(mapaMatricula map[uint]map[string]string) {
	for matricula, aluno := range mapaMatricula {
		fmt.Printf("Matricula: %d \n", matricula)

		for nome, curso := range aluno {
			fmt.Printf("Nome: %s. Curso: %s \n", nome, curso)
		}
	}
}

func mapaString(novoMapa map[string]string) {
	for nome, curso := range novoMapa {
		fmt.Println(nome, curso)
	}
}

func main() {
	fmt.Println("-- CRIAÇÃO DE MAPAS --")

	fruteira := map[string]float32{

		"Mamão":    5.99,
		"Banana":   3.50,
		"Maçã":     4.80,
		"Melancia": 12.00,
	}

	// Adicionando um novo elemento no mapa
	fruteira["Uva"] = 1.99

	//Fazendo a criação de um mapa dentro de outro mapa:
	alunosGraduacao := map[uint]map[string]string{
		202066974: {
			"Vinícius Cardoso": "Engenharia de computação",
		},

		21105576: {
			"Ana clara": "Engenharia de redes de comunicação",
		},
	}

	// Acessando os elementos que está presente no mapa. E acessando os elementos de um mapa dentro de outro mapa.
	leituraMapa(fruteira)
	fmt.Println()
	leituraMapaMapa(alunosGraduacao)

	// Para deletar um elemento dentro de um mapa.
	fmt.Println("Digite a matrícula do aluno que deseja deletar: ")
	var matricula uint

	fmt.Scan(&matricula)

	// Crio uma variável underscore para que o go não peça que eu use a variável. A variável existe me retorna true ou false para saber se o elemento se encontra no mapa.
	_, existe := alunosGraduacao[matricula]

	if existe {
		delete(alunosGraduacao, matricula)
		fmt.Printf("Aluno com a matrícula %d foi deletado com sucesso!\n\n", matricula)
	} else {
		fmt.Printf("Erro: A matrícula %d não foi encontrada no sistema.\n\n", matricula)
	}

	leituraMapaMapa(alunosGraduacao)

}
