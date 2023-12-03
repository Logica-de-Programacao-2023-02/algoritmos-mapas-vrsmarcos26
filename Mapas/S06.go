//Escreva uma função que receba uma lista de mapas, onde cada mapa contém a contagem de palavras de um texto,
//e retorne um único mapa contendo a soma de todas as contagens.

package main

import (
	"fmt"
)

func SomarContagens(listadeMapas []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, mapa := range listadeMapas {
		for palavra, contagem := range mapa {
			soma[palavra] += contagem
		}
	}

	return soma
}

func main() {
	mapa1 := map[string]int{"apple": 3, "banana": 2, "orange": 1}
	mapa2 := map[string]int{"apple": 2, "cherry": 5, "banana": 1}

	listaDeMapas := []map[string]int{mapa1, mapa2}

	resultado := SomarContagens(listaDeMapas)

	fmt.Println("Soma de contagens de palavras:")
	for palavra, contagem := range resultado {
		fmt.Printf("%s: %d\n", palavra, contagem)
	}
}
