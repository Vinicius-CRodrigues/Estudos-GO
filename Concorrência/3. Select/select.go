package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			canal1 <- "canal 1"
			time.Sleep(time.Second)
		}

	}()

	go func() {
		for {
			canal2 <- "canal 2"
			time.Sleep(time.Second * 4)
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)

		}
	}
}
