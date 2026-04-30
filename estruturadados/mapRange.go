// Range - Um Range é uma estrutura de dados que representa um intervalo de valores. Ele é comumente usado para armazenar e manipular intervalos de números, como intervalos de tempo, intervalos de índices, ou qualquer outro tipo de intervalo que possa ser representado por um par de valores (início e fim)

// utilizar um range dentro de um loop for para iterar sobre um map e imprimir suas chaves e valores

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
}