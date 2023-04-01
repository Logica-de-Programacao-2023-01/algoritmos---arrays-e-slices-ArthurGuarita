package main

import "fmt"

func main() {
	// Cria um slice de inteiros com tamanho 5
	numeros := make([]int, 5)

	// Solicita ao usuário que informe números inteiros até que o slice esteja cheio
	for i := 0; i < len(numeros); {
		var valor int
		fmt.Print("Digite um número inteiro: ")
		fmt.Scan(&valor)

		// Verifica se o número já está presente no slice
		encontrado := false
		for _, num := range numeros {
			if num == valor {
				encontrado = true
				break
			}
		}

		// Adiciona o número ao slice apenas se ele não estiver presente
		if !encontrado {
			numeros[i] = valor
			i++
		} else {
			fmt.Println("O número já está presente no slice.")
		}
	}

	// Imprime o slice resultante
	fmt.Println(numeros)
}
