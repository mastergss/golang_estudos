// Função para exibir mensagens, estáticas e dinâmicas, utilizando parâmetros e variáveis locais

package main

import "fmt"

func main() {
	mensagem := "Olá, mundo!"
	fmt.Println(mensagem)
	// chamar a função para exibir mensagem
	exibirMensagem()
	// chamar a função para exibir mensagem personalizada
	mensagemPersonalizada("Bem-vindo ao Go!")
}

// Função para exibir mensagem estática
func exibirMensagem() {
	fmt.Println("Exibir mensagem estática, Olá, mundo!")
}

// Função que recebe uma mensagem como parâmetro e a exibe
func mensagemPersonalizada(mensagem string) {
	fmt. Println("Mensagem Personalizada,", mensagem)
	mensagem += " Bom dia!"
	fmt.Println(mensagem)
}