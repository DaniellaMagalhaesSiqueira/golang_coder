package main

import "fmt"

func main() {
	numeros := [...]int{7, 2, 3, 4, 5} //compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // o _ é para ignorar o índice e colocar o valor de dentro do array
		fmt.Println(num)
	}
}
