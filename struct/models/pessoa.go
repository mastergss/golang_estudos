package models

import "time"

// Composição de structs: A struct Pessoa tem um campo Endereco do tipo Endereco, o que significa que cada pessoa tem um endereço associado a ela. Isso é um exemplo de composição, onde uma struct é composta por outra struct.
type Pessoa struct {
	Nome  		string
	//Idade		int
	Endereco	Endereco
	DataNascimento time.Time
}

// Método para calcular a idade de uma pessoa com base na data de nascimento. Ele recebe uma struct Pessoa como argumento e retorna a idade calculada
// O método Idade é definido para a struct Pessoa, permitindo que seja chamado diretamente em uma instância de Pessoa para obter a idade calculada com base na data de nascimento.
// Método é exclusivo para a struct Pessoa, enquanto a função CalcularIdade é uma função independente que pode ser usada para calcular a idade de qualquer pessoa, desde que seja passada como argumento.
func (p Pessoa) Idade() int {
	anoAtual := time.Now().Year()
	anoNascimento := p.DataNascimento.Year()
	idade := anoAtual - anoNascimento
	return idade
}

// Função para calcular a idade de uma pessoa com base na data de nascimento
func CalcularIdade(p Pessoa) int {
	anoAtual := time.Now().Year()
	anoNascimento := p.DataNascimento.Year()
	idade := anoAtual - anoNascimento
	return idade
}