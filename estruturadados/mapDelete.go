// Delete - Remover um elemento do mapa usando a função delete()

package main

import "fmt"

func main() {
	// criar um mapa utilizando a função make()
	cidade := make(map[string] int)
	// adicionar elementos ao mapa 'cidade'
	cidade["SP"] = 14000000
	cidade["RJ"] = 8000000
	cidade["BH"] = 6000000

	// iterar sobre o mapa utilizando range e imprimir chaves e valores
	for chave, valor := range cidade {
		fmt.Println("A cidade de", chave, "tem um total:", valor, "habitantes.")
	}

	// remover o elemento com a chave "RJ" do mapa utilizando a função delete()
	delete(cidade, "RJ")

	// iterar sobre o mapa para verificar se o elemento foi removido
	for chave, valor := range cidade {
		fmt.Println("A cidade de", chave, "tem um total:", valor, "habitantes.")
	}

	// verificar se a chave "RJ" ainda existe no mapa
	valor, existe := cidade["RJ"]
	if existe {
		fmt.Println("Valor da chave RJ: ", valor)
	} else {
		fmt.Println("A chave 'RJ' não foi encontrada no mapa 'cidade'")
	}
}