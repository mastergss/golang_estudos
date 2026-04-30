// Criar um array com 2 elementos do tipo inteiro, imprimir o array, o tamanho do array e a soma dos elementos utilizando um loop for e um loop range

package main

import "fmt"

func main() {
	lista := make([]int, 0)
	lista = append(lista, 10)
	lista = append(lista, 20)
	fmt.Println("Lista: ", lista)
	fmt.Println("Tamanho da lista: ", len(lista))

	somaTotal := 0
	for i:= 0; i < len(lista); i++ {
		somaTotal += lista[i]
	}
	fmt.Println("Soma Total da Lista: ", somaTotal)

	somaRange := 0
	for _, valor := range lista {
		somaRange += valor
	}
	fmt.Println("Soma Total utilizando Range: ", somaRange)
}