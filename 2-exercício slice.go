// Crie uma struct chamada Livro com os campos "título"
// ,
// "autor" e "ano de publicação".
// Escreva uma função que receba um Livro como parâmetro e imprima suas informações.
package main

import "fmt"

type Livro struct {
	título string
	autor  string
	ano    int
}

func pri(l Livro) {
	fmt.Printf("Título:%s\n", l.título)
	fmt.Printf("Autor:%s\n", l.autor)
	fmt.Printf("Ano:%s\n", l.ano)

}
func main() {
	var título, autor string
	var ano int
	fmt.Print("Escreva o nome titulo do livro")
	fmt.Scan(&título)
	fmt.Print("Escreva o nume do autor do livro")
	fmt.Scan(&autor)
	fmt.Print("Escreva o ano")
	fmt.Scan(&ano)
	recebe := Livro{
		título: título,
		autor:  autor,
		ano:    ano,
	}
	pri(recebe)
}
