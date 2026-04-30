// Ponteiro é uma referência para um valor na memória. Ele armazena o endereço de memória de uma variável, permitindo que você acesse e modifique o valor dessa variável indiretamente. Em Go, os ponteiros são usados para passar valores por referência, o que pode ser útil para otimizar o desempenho e evitar a cópia de grandes estruturas de dados

package main

import "fmt"

func main()  {
	x := 5
	// y := x -- anotação anterior, desta forma y recebe o valor de x, e qualquer alteração feita em y não refletirá em x, pois são variáveis independentes
	// Para apontar ao mesmo endereço de memória, é necessário usar o operador & para obter o endereço de memória da variável x e atribuí-lo a y, assim y se torna um ponteiro para x, e qualquer alteração feita em y refletirá em x, pois ambos apontam para o mesmo endereço de memória
	y := &x // Agora y é um ponteiro para x, e para acessar o valor de x através de y, é necessário usar o operador * para desreferenciar o ponteiro, ou seja, acessar o valor armazenado no endereço de memória apontado por y
	// y = 10 -- anotação anterior, desta forma y recebe o valor 10, e qualquer alteração feita em y não refletirá em x, pois são variáveis independentes
	*y = 10 // Agora, ao usar *y, estamos modificando o valor de x indiretamente através do ponteiro y, e qualquer alteração feita em y refletirá em x, pois ambos apontam para o mesmo endereço de memória
	fmt.Println("----- Função MAIN -----")
	// Imprime os valores de x e y, e os endereços de memória onde x e y estão armazenados
	fmt.Println("valor de x: ", x)
	fmt.Println("valor de y: ", y) // Como y é um ponteiro, devemos usar o operador * para acessar o valor apontado por y, ou seja, o valor de x
	fmt.Println("valor de y: ", *y)
	// A sintaxe & fornece o endereço de memória 
	fmt.Println("Endereço de memória: ", &x, &y) // Como y já aponta para o endereço de &x, devemos remover o &y para mostrar o endereço de memória de y, que é o mesmo de x, ou seja, &x
	fmt.Println("Endereço de memória: ", &x, y)

	// imprimirValores(x, y) -- anotação anterior, desta forma, a função imprimirValores recebe os valores de x e y, e qualquer alteração feita em x e y dentro da função não refletirá em x e y, pois são variáveis independentes
	// imprimirValores(x, *y) // Para chamar a função é necessário passar o valor apontado por *y
	imprimirValores(&x, y) // Agora para chamar a função devemos passar o endereço de memória &x e o ponteiro y, pois y já é um ponteiro para x
}

// func imprimirValores(x int, y int) { // Desta forma, é criada uma cópia dos valores x e y, e as alterações feitas não refletem em x e y, e aumentam o uso de memória, pois cada variável terá um novo espaço
func imprimirValores(x *int, y *int) { // Na função, devemos usar os ponteiros, ou seja, *int que indica que x e y são ponteiros inteiros
	fmt.Println("----- Função Imprimir Valores -----")
	fmt.Println("valor de x: ", x) 
	fmt.Println("valor de y: ", y)
	// Como x e y são ponteiros, devemos usar o operador * para acessar os valores apontados por x e y, ou seja, os valores de x e y
	fmt.Println("valor de x: ", *x)
	fmt.Println("valor de y: ", *y)
	fmt.Println("Endereço de memória: ", &x, &y)
	// Como x e y já são ponteiros, devemos remover o & para mostrar o endereço de memória apontado por x e y, que é o mesmo de x e y, ou seja, &x e &y
	fmt.Println("Endereço de memória: ", x, y)
	// Ao alterar o valor de *x, isso reflete em todas as variáveis que apontam para o mesmo endereço de memória, ou seja, x e y, pois ambos são ponteiros para o mesmo endereço de memória
	*x = 20
	fmt.Println("valor de x: ", *x)
	fmt.Println("valor de y: ", *y)
}