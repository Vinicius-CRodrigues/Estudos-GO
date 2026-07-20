package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Vinicius")
	for {
		fmt.Println(<-canal)
		time.Sleep(time.Second)
	}

}

// Crio uma função escrita que retorna um canal de strings
func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)

		}
	}()

	return canal
}
