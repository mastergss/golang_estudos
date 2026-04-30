package main

import "fmt"

func main() {
	x := 10
	//y := 5
	//fmt.Println("Valores de x: e y:", x, y)
	//fmt.Println("Endereço de memória de x: e y:", &x, &y)
	y := &x
	*y = 20
	fmt.Println("Valores de x: e y:", x, *y)
	fmt.Println("Endereço de memória de x: e y:", &x, y)

	imprimirValores(&x, y)
}

func imprimirValores(x *int, y *int) {
	*x = 30
	fmt.Println("Valor de x: ", *x, "Valor de y: ", *y)
	fmt.Println("End. meméria x: ", x, "End. memória y: ", y)
}