package workers

import "fmt"

func WorkerPool() {
	tarefas, respostas := make(chan int, 45), make(chan int, 45)

	for i := 0; i < 4; i++ {
		go Worker(tarefas, respostas)
	}

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resposta := <-respostas

		fmt.Println(resposta)
	}
}

// tarefa <-chan int: A seta está "saindo" do canal e indo para a esquerda. Isso significa que este é um canal de leitura. O Worker só pode puxar dados de dentro dele.
// Respostas chan<- int: A seta está apontando "para dentro" do canal. Isso significa que este é um canal de escrita. O Worker só pode enviar dados para ele.
/*
"Olá, eu sou o Worker. Eu olho para a esteira de tarefas e fico esperando um número chegar. Assim que chega um numero, eu calculo o fibonacci dele e coloco o resultado na esteira de respostas. Eu vou continuar fazendo isso até que desliguem a esteira de tarefas."
*/
func Worker(tarefa <-chan int, respostas chan<- int) {
	for numero := range tarefa {
		respostas <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-1) + fibonacci(posicao-2)

}
