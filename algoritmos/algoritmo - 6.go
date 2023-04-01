package main

import "fmt"

func main() {
	var numeros = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var valor int
	var achou bool
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for _, num := range numeros {
		if valor == num {
			achou = true
			break
		}
	}
	if achou {
		fmt.Println("O valor digitado foi encontrado.")
	} else {
		fmt.Println("Não foi encontrado.")
	}
}
