// Exemplo de como criar e utilizar lista

package main

import "fmt"

func main() {
	// lista estática
	lista := []int {1, 4, 6, 9}
	fmt.Println("Lista:", lista)
	fmt.Println("Posição 1:", lista[0])
	fmt.Println("Posição 2:", lista[1])

	fmt.Println("Tamanho da lista", len(lista))

	for i:= 0; i < len(lista); i++ {
		fmt.Println("Posição", i, ":", lista[i])
	}

	// adicionar valores
	lista = append(lista, 8)
	fmt.Println("Lista atualizada: ", lista)

	// criar lista
	lista2 := make([]int, 1)
		// adicionar valores
		lista2[0] = 5
		lista2 = append(lista2, 3)
	fmt.Println("Lista2: ", lista2)
	fmt.Printf("%T\n", lista2)

	// calcular média das listas
	// média lista
	total := 0
	for i:= 0; i < len(lista); i++ {
		total += lista[i]
	}
	media := total / len(lista)
	fmt.Println("Média da lista: ", media)
	// media lista2
	total2 := 0
	for i:= 0; i < len(lista2); i++ {
		total2 += lista2[i]
	}
	media2 := total2 / len(lista2)
	fmt.Println("Média da lista2: ", media2)
}