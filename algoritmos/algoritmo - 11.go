package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n [2][3]int
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[i]); j++ {
			n[i][j] = rand.Intn(100)
		}
	}
	//perguntar ao usuario
	var linha, coluna int
	fmt.Println("Digite a linha a ser exibido: ")
	fmt.Scan(&linha)
	fmt.Println("Digite a coluna a ser exibida: ")
	fmt.Scan(&coluna)

	if linha >= 0 && linha < len(n) && coluna >= 0 && coluna < len(n[0]) {
		fmt.Println(n[linha][coluna])
	} else {
		fmt.Println("Índices inválidos")
	}
}
