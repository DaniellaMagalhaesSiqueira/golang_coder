package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15450.78,
		"Pedro Junior":   1750.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Inexistente")

	fmt.Println(funcsESalarios)
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
