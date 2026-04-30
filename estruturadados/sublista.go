// criar uma sublista e exibir números menores que 5

package main

import "fmt"

func main() {
	listaCompleta := []int {2, 6, 3, 9, 8, 4, 1}
	listaMenorCinco := make([]int, 0)
	for i:= 0; i < len(listaCompleta); i++ {
		if listaCompleta[i] < 5 {
			listaMenorCinco = append(listaMenorCinco, listaCompleta[i])
			fmt.Println("Posição [",i,"]", listaCompleta[i])
			fmt.Println("Lista Menor que Cinco:", listaMenorCinco)
		}
	}

	fmt.Println(listaMenorCinco)
}