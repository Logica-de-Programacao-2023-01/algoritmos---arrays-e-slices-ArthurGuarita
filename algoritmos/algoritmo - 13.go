package main

import "fmt"

func main() {
	var array [7]int
	//solicitar ao usuario
	var valor int
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	//adição
	array[0], array[6] = valor, valor
	fmt.Println(array)
}
