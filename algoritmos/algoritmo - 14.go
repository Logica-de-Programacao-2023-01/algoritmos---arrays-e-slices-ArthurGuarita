package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slime = make([]int, 8)
	for i := 0; i < len(slime); i++ {
		slime[i] = rand.Intn(100)
	}
	fmt.Println(slime)
	//indices a serem trocados
	var indice1, indice2, invert int
	fmt.Println("Digite o primeiro indice: ")
	fmt.Scan(&indice1)
	fmt.Println("Digite o segundo indice: ")
	fmt.Scan(&indice2)
	invert = slime[indice1]
	slime[indice1] = slime[indice2]
	slime[indice2] = invert

	fmt.Println(slime)
}
