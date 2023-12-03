//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os valores são as quantias de
//dinheiro que cada pessoa gastou em uma conta compartilhada. A função deve calcular o valor que cada pessoa
//deve receber ou pagar para igualar as despesas.

package main

import (
	"fmt"
)

func EqualizarDespesas(gastos map[string]float64) map[string]float64 {
	totalGasto := 0.0

	for _, valor := range gastos {
		totalGasto += valor
	}

	media := totalGasto / float64(len(gastos))

	saldos := make(map[string]float64)
	for pessoa, valor := range gastos {
		saldos[pessoa] = valor - media
	}

	return saldos
}

func main() {
	gastos := map[string]float64{"Alice": 30.0, "Bob": 20.0, "Charlie": 10.0}

	saldos := EqualizarDespesas(gastos)

	fmt.Println("Saldos para equalizar as despesas:")
	for pessoa, saldo := range saldos {
		fmt.Printf("%s: %.2f\n", pessoa, saldo)
	}
}
