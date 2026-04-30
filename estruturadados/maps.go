// map - chave e valor
// Mapas são uma estrutura de dados integrada, conveniente e poderosa, que associa valores de um tipo (a chave ) a valores de outro tipo (o elemento ou valor). O tipo de chave deve ser comparável; o tipo de valor pode ser qualquer tipo. Mapas são usados para representar associações, como a associação entre nomes e números de telefone em uma agenda telefônica

package main

import "fmt"

func main() {
	// criar um mapa utilizando a sintaxe literal e adicionar elementos
	mod := map[string] int {"SP": 10000000, "CG": 900000}
	fmt.Println(mod)

	// criar um mapa utilizando a função make()
	mod1 := make(map[string] int)
	// adicionar elementos ao mapa
	mod1["RJ"] = 8000000
	mod1["MG"] = 7000000
	fmt.Println(mod1)

	// acessar elementos do mapa utilizando a chave
	fmt.Println("Valor da chave SP:", mod["SP"])

	// verificar se uma chave existe no mapa
	valor, existe := mod["RJ"]
	if existe {
		fmt.Println("Valor da chave RJ: ", valor)
	} else {
		fmt.Println("Chave RJ não existe no Mapa 'mod'")
	}
}