package main

import "fmt"

func main() {
	var n [6]float64
	//Usuario informa valor
	var valor float64
	fmt.Print("Informe um número a ser adicionado ao array: ")
	fmt.Scan(&valor)
	//Adicionando o valor
	for i := 0; i < len(n); i++ {
		n[i] = valor
	}
	fmt.Println(n)
}
