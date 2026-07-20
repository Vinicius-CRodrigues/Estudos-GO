package main

import (
	"ferramenta-testes/address"
	"fmt"
)

func main() {
	fmt.Println(address.TypeOfAddresses("Maria dos milagres"))

	// Como estou usando uma struct de fora, preciso deixar claro qual variável estou preenchendo para caso haja alterações futuras, o struct não quebre.
}
