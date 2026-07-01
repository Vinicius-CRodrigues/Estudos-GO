package main

import (
	"fmt"
)

func main() {
	palavra := "-- LOOPS --"
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}

	var (
		nome     string
		idade    int
		opcao    string
		qdeNotas int
		nota     int
	)

	listaAlunos := make(map[string]int)
	notasAlunos := make([]int, 0, 100)

	for {

		fmt.Println()
		fmt.Println("1. Inserir nome e aluno: ")
		fmt.Println("2. Listar alunos e matricula: ")
		fmt.Println("3. Inserir notas: ")
		fmt.Println("4. Listar média das notas: ")
		fmt.Println("5. SAIR ")

		fmt.Scan(&opcao)

		if opcao == "5" {
			fmt.Println("FIM DO PROGRAMA !")
			break
		}

		switch opcao {
		case "1":
			fmt.Print("Insira o nome: ")
			fmt.Scan(&nome)
			fmt.Print("Insira o idade: ")
			fmt.Scan(&idade)

			listaAlunos[nome] = idade

		case "2":
			for nome, idade := range listaAlunos {
				fmt.Printf("Nome: %s. Idade: %d.", nome, idade)
			}

		case "3":
			fmt.Print("Insira a quantidade de notas que deseja adicionar: ")
			fmt.Scan(&qdeNotas)

			for i := 0; i < qdeNotas; i++ {
				fmt.Printf("Nota %d: ", i+1)
				fmt.Scan(&nota)

				notasAlunos = append(notasAlunos, nota)
			}

		case "4":
			soma := 0

			for i := range notasAlunos {
				soma += notasAlunos[i]
			}

			var media float64
			if len(notasAlunos) > 0 {
				media = float64(soma) / float64(len(notasAlunos))
			} else {
				media = 0
			}

			fmt.Printf("A média da turma é: %.2f\n", media)
		default:
			fmt.Println("Opção invalida !")

		}

	}

}
