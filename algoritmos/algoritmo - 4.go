package main

import "fmt"

func main() {
	var numeros []int
	var t, valor, result int
	fmt.Println("Digite o tamanho do slice: ")
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&valor)
		numeros = append(numeros, valor)
	}
	for _, num := range numeros {
		result += num
	}
	fmt.Println("A soma é", result)
	fmt.Println("O slice é ", numeros)
}
