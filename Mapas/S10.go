//Escreva uma função que receba um slice de inteiros e retorne um mapa onde as chaves são os pares de números
//encontrados no slice e os valores são as frequências de cada par.

package main

import "fmt"

func FrequenciaPares(nums []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pair := [2]int{nums[i], nums[j]}
			frequencia[pair]++
		}
	}

	return frequencia
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 1, 2}

	resultado := FrequenciaPares(numeros)

	fmt.Println("Frequência de pares de números:")
	for par, freq := range resultado {
		fmt.Printf("%v: %d\n", par, freq)
	}
}
