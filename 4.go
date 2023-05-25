// Crie uma struct chamada Playlist com os campos "nome" e "músicas".
// O campo "músicas" deve ser um slice de structs, cada uma representando uma música com os campos "título", "artista" e "duração".
// Escreva uma função que receba uma Playlist como parâmetro
// e imprima o nome da playlist, o tempo total da playlist e as informações de cada música.
package main

import "fmt"

type música struct {
	título  string
	artista string
	duração int
}
type Playlist struct {
	nome   string
	música []música
}

func play(p Playlist) {
	fmt.Printf("Nome %s", p.nome)

	for _, r := range p.música {
		fmt.Printf("título:%s\n", r.título)
		fmt.Printf("artista:%s\n", r.artista)
		fmt.Printf("duração:%d\n", r.duração)
	}
}
func main() {
	a := Playlist{
		nome: "Go",
		música: []música{
			{título: "m1", artista: "m2", duração: 2},
			{título: "j", artista: "peixe", duração: 10},
		},
	}
	play(a)
}
