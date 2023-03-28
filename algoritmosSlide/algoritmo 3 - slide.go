package main

import "fmt"

func main() {
	var valores = []int{1, 2, 3, 4, 5}
	valores = append(valores[:2], valores[3:]...)
	fmt.Println(valores)
}
