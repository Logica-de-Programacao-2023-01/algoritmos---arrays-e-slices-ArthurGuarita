package main

import "fmt"

func main() {

	//Criando o slice e atribuindo valores
	var n = make([]string, 8)
	n[0] = "Banana"
	n[1] = "Uva"
	n[2] = "Pêssego"
	n[3] = "Manga"
	n[4] = "Blueberry"
	n[5] = "Cereja"
	n[6] = "Acerola"
	n[7] = "Maçã"
	//Valor digitado pelo usuário

	var valor string
	fmt.Print("Digite um tipo de fruta: ")
	fmt.Scan(&valor)
	//Verificando se está presente no slice
	for i := 0; i < len(n); i++ {
		if n[i] == valor {
			n = append(n[:i], n[i+1:]...)
			i--
		}
	}
	fmt.Println(n)
}
