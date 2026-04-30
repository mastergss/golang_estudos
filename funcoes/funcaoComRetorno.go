// Função com retorno - utilizar a instrução 'return' para retornar um valor a partir de uma função que a chama

package main

import "fmt"

func main() {
	resultado := somar(5, 3)
	fmt.Println("Resultado: ", resultado)
}

// Função que recebe dois números inteiros, realiza a soma e retorna o resultado
func somar(num1 int, num2 int) int {
	resultado := num1 + num2
	return resultado
}