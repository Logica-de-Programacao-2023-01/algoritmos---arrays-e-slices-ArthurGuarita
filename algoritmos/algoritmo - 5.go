package main

import (
	"fmt"
)

func main() {
	var matriz [3][2]int
	for linha := range matriz {
		for coluna := range matriz[linha] {
			var value int
			fmt.Print("Digite o valor: ")
			fmt.Scan(&value)
			matriz[linha][coluna] = value
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Println(matriz[i])
	}
}
