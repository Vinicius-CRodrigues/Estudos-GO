package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("-- ESTRUTURA DE CONTROLE --")

	var node *int

	numero := 10

	node = &numero

	if novoNumero := *node; novoNumero == 10 {
		fmt.Println("O nó está pegando o valor de numero, e o novoNumero tem o valor ", novoNumero)
	}

	nomeIdade := make(map[string]int)

	nomeIdade["Vinicius"] = 24

	nomeIdade["Adileuza"] = 47

	nomeIdade["Ana Clara"] = 22

	// Consigo declarar uma variável dentro do if construindo ela com ponto e virgula no final e atribuir a condição.
	var parametro string
	fmt.Scanf("%s", &parametro)

	// A variável fica limitada ao escopo if init.
	if nome, existeNoMap := nomeIdade[parametro]; existeNoMap {
		fmt.Println(nome)
		fmt.Println(reflect.TypeOf(existeNoMap))

	}

}
