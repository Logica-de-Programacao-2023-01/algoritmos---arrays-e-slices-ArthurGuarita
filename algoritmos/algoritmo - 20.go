package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)
	//
	var array = make([]int, n)
	fmt.Println("Digite valores para o array: ")
	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
	}
	//verificação
	var crescente bool
	crescente = true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			crescente = false
			break
		}
	}
	if crescente {
		fmt.Println("O array está em ordem crescente.")
	} else {
		fmt.Println("O array não está em ordem crescente.")
	}
}
