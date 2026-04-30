// A função init é uma função especial em Go que é executada automaticamente antes da função main. Ela é usada para inicializar variáveis, configurar o ambiente ou realizar qualquer configuração necessária antes do programa começar a executar a lógica principal. A função init é chamada apenas uma vez, e pode ser definida em qualquer arquivo do pacote, mas deve ser declarada sem parâmetros e sem valor de retorno

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

func init() {
	fmt.Println("Função init, sendo iniciada antes de qualquer função, inclusive a função main")
}