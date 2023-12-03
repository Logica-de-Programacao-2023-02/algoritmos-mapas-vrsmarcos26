//Escreva uma função que gere a sequência de Fibonacci até um determinado número n e retorne um mapa onde as
//chaves são os índices da sequência e os valores são os números correspondentes.

package main

import "fmt"

func FibonacciAteN(n int) map[int]int {
	fibonacci := make(map[int]int)

	a, b := 0, 1
	index := 0

	for a <= n {
		fibonacci[index] = a
		index++
		a, b = b, a+b
	}

	return fibonacci
}

func main() {
	n := 50
	resultado := FibonacciAteN(n)

	fmt.Println("Sequência de Fibonacci até", n, ":")
	for indice, valor := range resultado {
		fmt.Printf("%d: %d\n", indice, valor)
	}
}
