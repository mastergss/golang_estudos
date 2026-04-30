// Função com retornos nomeados - é possível nomear e definir os tipos de retornos na assinatura do método, e assim basta utilizar a instrução 'return' sem especificar os valores, pois os valores já foram definidos na função, e o retorno será realizado de acordo com a ordem dos valores definidos na assinatura do método

package main

import "fmt"

func main() {
	somar, subtrair, dividir, multiplicar := operacoes(3, 2)
	fmt. Println("Somar: ", somar)
	fmt. Println("Subtrair: ", subtrair)
	fmt. Println("Dividir: ", dividir)
	fmt. Println("Multiplicar: ", multiplicar)
}

func operacoes(num1 int, num2 int) (somar int, subtrair int, dividir float64, multiplicar int) {
	somar = num1 + num2
	subtrair = num1 - num2
	dividir = float64(num1) / float64(num2)
	multiplicar = num1 * num2
	return
}

/*
func operacoes(num1 int, num2 int) (int, int, float64, int) {
	somar := num1 + num2
	subtrair := num1 - num2
	dividir := float64(num1) / float64(num2)
	multiplicar := num1 * num2
	return somar, subtrair, dividir, multiplicar
}
*/