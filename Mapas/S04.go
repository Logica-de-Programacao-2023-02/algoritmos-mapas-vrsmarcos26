//Escreva uma função que receba um slice de palavras e retorne um mapa onde as chaves são palavras que são
//anagramas entre si e os valores são os grupos de palavras anagramas.

package main

import "fmt"

func SomarValoresMapa(mapa map[string]int) int {
	soma := 0
	for _, valor := range mapa {
		soma += valor
	}
	return soma
}

func main() {
	valores := map[string]int{"a": 10, "b": 20, "c": 30}

	resultado := SomarValoresMapa(valores)

	fmt.Println("Soma dos valores do mapa:", resultado)
}
