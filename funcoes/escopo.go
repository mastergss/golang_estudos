// Escopo de variáveis em Go
// O escopo de uma variável é a região do código onde ela é acessível. Em Go, existem dois tipos de escopo: global e local. Variáveis globais são declaradas fora de qualquer função ou método e podem ser acessadas em todo o código. Variáveis locais são declaradas dentro de uma função ou método e só podem ser acessadas dentro desse bloco de código. No exemplo abaixo, a variável numMagico é declarada como uma variável global, enquanto a variável numMagico dentro do método calcular é uma variável local, acessível apenas dentro desse método.

package main

import "fmt"

var numMagico = 8 // variável com escopo global, acessível em todo o código, inclusive dentro do método operacoes

func main() {
	somar, subtrair, dividir, multiplicar := calcular(3, 2)
	fmt. Println("Somar: ", somar)
	fmt. Println("Somar: ", somar + numMagico) // variável numMagico não é acessível no escopo do main, pois é uma variável local do método operacoes // para acessar a variável numMagico, é necessário declará-la como uma variável global, ou seja, fora de qualquer função ou método
	fmt. Println("Subtrair: ", subtrair)
	fmt. Println("Dividir: ", dividir)
	fmt. Println("Multiplicar: ", multiplicar)
}

func calcular(num1 int, num2 int) (somar int, subtrair int, dividir float64, multiplicar int) {
	numMagico := 1 // variável com escopo local, acessível apenas dentro do método operacoes
	fmt.Println("Número Mágico: ", numMagico)
	// somar = num1 + num2
	somar = num1 + num2 + numMagico
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