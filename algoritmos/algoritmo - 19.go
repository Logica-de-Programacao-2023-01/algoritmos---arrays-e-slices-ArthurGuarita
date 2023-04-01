package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&n)

	//array 1

	var array1 = make([]int, n)
	fmt.Println("Digite os valores para o primeiro array: ")
	for i := 0; i < len(array1); i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println(array1)

	//array 2

	var array2 = make([]int, n)
	fmt.Println("Digite valores para o segundo array: ")
	for i := 0; i < len(array2); i++ {
		fmt.Scan(&array2[i])
	}
	// soma
	var array3 = make([]int, n)
	for i := 0; i < len(array3); i++ {
		array3[i] = array1[i] + array2[i]
	}
	fmt.Println(array3)
}
