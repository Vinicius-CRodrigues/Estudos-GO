package main

import "fmt"

type atacante interface {
	atacar() string
}

type guerreiro struct {
	nome  string
	forca int
}

type mago struct {
	nome       string
	forcaMagia int
}

func (g guerreiro) atacar() string {
	dano := 2 * g.forca
	return fmt.Sprintf("%s atacou com a força %d", g.nome, dano)

}

func (m mago) atacar() string {
	danoMago := 2 * m.forcaMagia

	return fmt.Sprintf("%s atacou com a força %d", m.nome, danoMago)

}

func descreverPersonagem(atack atacante) {
	fmt.Println(atack.atacar())

}

func main() {
	userGuerreiro := guerreiro{"Barbaro", 10}

	userMago := mago{"Veigar", 1000}

	descreverPersonagem(userGuerreiro)
	descreverPersonagem(userMago)

}
