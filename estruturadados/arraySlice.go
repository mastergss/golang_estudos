// Array e Slice - são estruturas de dados que armazenam uma coleção de elementos do mesmo tipo. A diferença entre eles é que o array tem um tamanho fixo, enquanto o slice é dinâmico e pode crescer ou diminuir de tamanho.

package main

import "fmt"

func main() {
	// slice é uma estrutura de dados dinâmica, pode crescer ou diminuir de tamanho, diferente do array que tem um tamanho fixo
	listaSlice := []int {1, 4, 6, 9, 2, 10, 5, 3}
	fmt.Println("Lista completa: ", listaSlice)
	fmt.Printf("%T \n", listaSlice)
	listaSlice = append(listaSlice, 11)
	// ao adicionar um novo elemento, uma nova lista é criada, e a antiga é descartada, isso influcencia na performance
	fmt.Println("Nova lista: ", listaSlice)

	listaArray := [8]int {1, 4, 6, 9, 2, 10, 5, 3}
	fmt.Println("Lista completa: ", listaArray)
	fmt.Printf("%T \n", listaArray)
	// ao adicionar um novo elemento, a lista é modificada, e não é necessário criar uma nova lista, isso melhora a performance
	// listaArray = append(listaArray, 11) // isso não é possível, pois o array tem um tamanho fixo
}