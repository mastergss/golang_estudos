// Criar um array com 9 elementos ddo tipo inteiro, imprimir o array, o tamanho do array e a soma dos elementos utilizando um loop for e um loop range, separando a soma dos números de 1 a 5 e de 6 a 10

package main

import "fmt"

func main() {
	lista := []int {2, 8, 3, 10, 5, 4, 7, 9, 1}
	somarUmAteCinco := 0
	somarSeisAteDez := 0

	for i:= 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			somarUmAteCinco += lista[i]
		} else {
			somarSeisAteDez += lista[i]
		}
	}
	fmt.Println("A soma da lista dos números de 1 a 5 = ", somarUmAteCinco)
	fmt.Println("A soma da lista dos números de 6 a 10 = ", somarSeisAteDez)

	listaAteCinco := 0
	listaAteDez := 0

	for _, valor := range lista {
		if valor <= 5 {
			listaAteCinco += valor
		} else {
			listaAteDez += valor
		}
	}
	fmt.Println("A soma range da lista até 5 = ", listaAteCinco)
	fmt.Println("A soma range da lista até 10 = ", listaAteDez)
}