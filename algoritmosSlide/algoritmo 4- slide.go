package main

import "fmt"

func main() {
	var numeros = [6]float64{1, 2, 3, 4, 5, 6}
	var result float64
	result = 0
	for _, num := range numeros {
		result = result + num
	}
	var media float64
	media = result / 6
	fmt.Print(media)
}
