package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// declarar e inicializar de uma forma reduzida
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)
	//Em GO dá erro de compilação caso a variável é definida e não utilizada
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	//GO é uma linguagem fortemente tipada, os tipos não variam, ou seja, g sempre será um int
	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)

}
