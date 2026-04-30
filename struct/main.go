// Structs são tipos de dados compostos que agrupam variáveis relacionadas sob um mesmo nome. Elas são usadas para representar objetos do mundo real, como pessoas, carros, endereços, etc. Cada campo em uma struct pode ter um tipo diferente, e as structs podem ser aninhadas dentro de outras structs.
package main

import (
	"fmt"
	"golangestudo/models"
	"time"
)

/* mover para a pasta models
uma struct com letra minúscula é privada, ou seja, só pode ser acessada dentro do mesmo pacote. Já uma struct com letra maiúscula é pública, podendo ser acessada de outros pacotes.
type endereco struct {
	rua		string
	numero	int
	cidade	string
}
*/

func main() {
	fmt.Println("Iniciando...")
	/*
		var end1 endereco
		end1.rua = "Rua dos Bobos"
		end1.numero = 0
		end1.cidade = "São Paulo"
		fmt.Println(end1)

		end2 := endereco{
			rua: "Rua 1",
			numero: 2,
			cidade: "Rio de Janeiro",
		}
		fmt.Println(end2)
	*/

	var end1 models.Endereco
	end1.Rua = "Rua dos Bobos"
	end1.Numero = 0
	end1.Cidade = "São Paulo"
	fmt.Println(end1)

	end2 := models.Endereco{
		Rua: "Rua Um",
		Numero: 2,
		Cidade: "Rio de Janeiro",
	}
	fmt.Println(end2)

	var pes1 models.Pessoa
	pes1.Nome = "Gerson"
	//pes1.Idade = 30
	pes1.Endereco = end1
	pes1.DataNascimento = time.Date(1981, 7, 31, 0,0,0,0, time.UTC)
	fmt.Println(pes1)
	// função para calcular a idade de uma pessoa com base na data de nascimento
	calcularIdade := models.CalcularIdade(pes1)
	fmt.Println(calcularIdade)
	// método para calcular a idade de uma pessoa com base na data de nascimento
	idade := pes1.Idade()
	fmt.Println(idade)

	pes2 := models.Pessoa{
		Nome: "Maria",
		//Idade: 25,
		Endereco: end2,
		DataNascimento: time.Date(1996, 5, 15, 0,0,0,0, time.UTC),
	}
	fmt.Println(pes2)
	// função para calcular a idade de uma pessoa com base na data de nascimento
	calcularIdade = models.CalcularIdade(pes2)
	fmt.Println(calcularIdade)
	// método para calcular a idade de uma pessoa com base na data de nascimento
	idade = pes2.Idade()
	fmt.Println(idade)
}
