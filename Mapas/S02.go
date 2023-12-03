//Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos os elementos dos mapas de entrada.
//Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.

package main

import "fmt"

func MesclarMapas(mapa1, mapa2 map[string]string) map[string]string {
	resultado := make(map[string]string)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	mapa2 := map[string]string{"b": "blueberry", "d": "date", "e": "elderberry"}

	resultado := MesclarMapas(mapa1, mapa2)

	fmt.Println("Resultado da mesclagem de mapas:")
	for chave, valor := range resultado {
		fmt.Printf("%s: %s\n", chave, valor)
	}
}
