package models

// A struct Endereco representa um endereço com os campos Rua, Numero e Cidade. Ela é exportada (pública) porque começa com letra maiúscula, permitindo que seja acessada de outros pacotes.
type Endereco struct {
	Rua		string
	Numero	int
	Cidade	string
}