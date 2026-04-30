// utilizar index para acessar e exibir elementos de uma lista

package main

import "fmt"

func main() {
	lista := []int {1, 4, 6, 9, 2, 10, 5, 3}
	fmt.Println("Lista completa: ", lista)

	// acessar os elementos utilizando index, acessar os 3 primeiros elementos
	segundaLista := lista[:3]
	fmt.Println("Segunda Lista: ", segundaLista)
	// acessar e exibir os elementos a partir do index 4
	terceiraLista := lista[4:]
	fmt.Println("Terceira Lista: ", terceiraLista)
	// acessar e exibir o último elemento da lista
	ultimoItem := lista[len(lista)-1]
	fmt.Println("Último ítem: ", ultimoItem)
}