// Função com múltiplos retornos - é possível retornar mais de um valor a partir de uma função, utilizando a instrução 'return' e definindo os tipos na assinatura do método

package main

import "fmt"

func main() {
	soma, subtracao, divisao, multiplicacao := operacoesMatematicas(8, 4)
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtracao)
	fmt.Println("Divisão: ", divisao)
	fmt.Println("Multiplicação: ", multiplicacao)
}
// Função que recebe dois números inteiros e retorna as operações básicas da matermática
// quando há mais de um retorno, é necessário definir o tipo na assinatura do método, utilizando parênteses
// func operacoesMatematicas(num1 int, num2 int) (int, int, float64, int) {
func operacoesMatematicas(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subtracao := num1 - num2
	// divisao := float64(num1) / float64(num2) // para evitar divisão inteira, é necessário converter os números para float64
	divisao := num1 / num2
	multiplicacao := num1 * num2
	// a conversão para float64 pode ser realizada no return, e assim converter o resultado da divisão para exibir o resultado com as casas decimais
	// return soma, subtracao, float64(divisao), multiplicacao

	// para retornar múltiplos valores, basta colocar em ordem separados por vírgula
	return soma, subtracao, divisao, multiplicacao
}