package main

import "fmt"

func main() {
	var numeros = [4]float64{1, 2, 3, 4}
	var result float64
	result = 1
	for _, num := range numeros {
		result = result * num
	}
	fmt.Println(result)
}
