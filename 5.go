//Usando as mesmas structs do exercício anterior, escreva uma função que receba um título de música como parâmetro
//e retorne uma lista das Playlists que possuem esse título.
package main

type Musica struct {
	Titulo  string
	Artista string
	Duracao int
}

type playlist struct {
	Nome    string
	Musicas []Musica
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {

}
