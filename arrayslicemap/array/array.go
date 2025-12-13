package main

import "fmt"

func main() {
	//homogênea e estática (se tiver 10 posições vai ficar com 10)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.6, 7.3
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
