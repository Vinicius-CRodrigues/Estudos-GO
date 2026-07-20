package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Crio uma variável que represeta um grupo de espera.
	var waitGroup sync.WaitGroup

	// Informo que a variável terá dois grupos de espera na fila.
	waitGroup.Add(2)

	// Criamos funções anonimas em goRoutines.
	go func() {
		escreverNaTela("Teste_A")
		waitGroup.Done() // Retira -1 do contador de rotinas na fila de espera.

	}()

	go func() {
		escreverNaTela("Teste_B")
		waitGroup.Done()

	}()

	waitGroup.Wait() // Esperar a contagem das goRoutines chegar em zero.
}

func escreverNaTela(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto, i+1)
		time.Sleep(time.Second)
	}
}
