package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var array [10]float64
	var n []float64

	// randomizar elementos do array
	for i := 0; i < len(array); i++ {
		randomNum := rand.Float64() * 100
		rounded := math.Round(randomNum*100) / 100 // arredondar para 2 casas decimais
		array[i] = rounded
	}

	// verificação
	for _, num := range array {
		if num > 5 {
			n = append(n, num)
		}
	}

	// imprimir resultado formatado
	fmt.Printf("Números arredondados maiores que 5:\n")
	for _, num := range n {
		fmt.Printf("%.2f | ", num)
	}
}
