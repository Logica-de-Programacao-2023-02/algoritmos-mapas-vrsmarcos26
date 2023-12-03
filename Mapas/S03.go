//Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

package main

import "fmt"

func somarValores(m map[int]int) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

func main() {
	meuMapa := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}

	resultado := somarValores(meuMapa)
	fmt.Printf("A soma dos valores é: %d\n", resultado)
}

