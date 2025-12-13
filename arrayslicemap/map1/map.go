package main

import "fmt"

func main() {
	// var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[123456789] = "Maria"
	aprovados[123456781] = "Pedro"
	aprovados[123456748] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[123456748])

	delete(aprovados, 123456781)
	fmt.Println(aprovados[123456781])
}
