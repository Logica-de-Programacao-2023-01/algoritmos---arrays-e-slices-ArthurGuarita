package main

import (
	"fmt"
)

func main() {
	var matriz [3][4]int
	for linha := range matriz {
		for coluna, elemento := range matriz[linha] {
			elemento = linha + coluna
			fmt.Println("ELemento[", linha, "][", coluna, "] = ", elemento)
		}
	}
}
