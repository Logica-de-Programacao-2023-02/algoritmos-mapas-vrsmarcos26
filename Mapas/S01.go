//Escreva uma função que conte a ocorrência de cada palavra em uma string e retorne um mapa onde as chaves
//são as palavras encontradas e os valores são o número de ocorrências de cada palavra.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ContarOcorrenciasPalavras(frase string) map[string]int {
	palavras := strings.Fields(frase) // Separa a string em palavras

	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		palavra = limparPalavra(palavra)
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func limparPalavra(palavra string) string {
	regexpLimpeza := regexp.MustCompile(`[[:punct:]]`)
	return strings.ToLower(regexpLimpeza.ReplaceAllString(palavra, ""))
}

func main() {
	frase := "Esta é uma frase de exemplo. Esta frase contém palavras repetidas. Exemplo de frase."
	resultado := ContarOcorrenciasPalavras(frase)

	fmt.Println("Ocorrências de palavras:")
	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
