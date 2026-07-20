package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Cliente struct {
	Nome  string `json:"nome"` // Utilizando TAGS para assinalar JSON
	Idade uint   `json:"idade"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

type RacasDeCavalo struct {
	ID           int      `json:"id"`
	Nome         string   `json:"nome"`
	Origem       string   `json:"origem"`
	Aptidao      []string `json:"aptidao"`
	Temperamento string   `json:"temperamento"`
}

func main() {

	clienteID := Cliente{Nome: "Maria", Idade: 70, Email: "maria@gmail.com", CPF: 123}

	// Marshal cria um json com as informações do struct.
	clienteIDJSON, err := json.Marshal(clienteID)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Quando eu fizer o print, ela me dará um slice e bytes que representa os JSON formatados.
	fmt.Println(clienteIDJSON)

	//Para eu conseguir fazer o print do JSON, eu uso a seguinte sintaxe.
	fmt.Println(string(clienteIDJSON))

	//Também posso criar json de mapas da seguinte forma
	mapa := make(map[string]string)

	mapa["Nome"] = "Vinicius"
	mapa["Curso"] = "Engenharia de computação"
	mapa["Modalidade"] = "Graduação"

	fmt.Println(mapa)

	mapaJSON, err := json.Marshal(mapa)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(mapaJSON))

	//Lendo um arquivo json com o pacote os:
	testeJSON, err := os.ReadFile("teste.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	// Criando um slice com a struct RacasDeCavalo
	dadosRaca := []RacasDeCavalo{}

	// Transformando o json de um arquivo externo em uma struct.
	json.Unmarshal([]byte(testeJSON), &dadosRaca)

	fmt.Println("---------------------------JSON--------------------------------")

	fmt.Printf("%+v\n", dadosRaca)

	fmt.Println("---------------------------JSON--------------------------------")

	// Percorrendo pelo novo struct com o for.
	for _, raça := range dadosRaca {
		fmt.Println("Raça:", raça.Nome)
		fmt.Println("Origem:", raça.Origem)
		for _, habilidade := range raça.Aptidao {
			fmt.Println("Habilidades da raça:", habilidade)
		}
	}

}
