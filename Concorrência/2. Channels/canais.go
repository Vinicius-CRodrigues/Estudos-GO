package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	// Criando uma goRoutine com a função para que ela possa ser executada em conjunto com o canal.
	go escreverNaTela("texto", canal)

	// Para cada mensagem recebida no canal enquanto ele estiver aberto, eu vou printar o conteúdo que está em canal.
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	/*
		// Também posso construir de uma forma que vou conferindo se o canal ainda permanece aberto da seguinte forma:
		for {
		// Rcebendo os dados no canal.
			mensagem, aberto := <- canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/
	fmt.Println("Fim do programa !")

}

func escreverNaTela(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// Enviando a informações texto para dentro do canal.
		canal <- texto
		time.Sleep(time.Second)
	}
	// Após o loop executar por 5 vezes, ele não irá mais receber dados no canal.
	close(canal)
}
