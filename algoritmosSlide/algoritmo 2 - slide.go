package main

import "fmt"

func main() {
	var nomes []string
	nomes = append(nomes, "João", "Maria", "José")
	for i := 0; i < 3; i++ {
		fmt.Println(nomes[i])
	}
}
