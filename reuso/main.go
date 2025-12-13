package main

import (
	"fmt"

	"cod3r2/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// fmt.Println(area.TrianguloEq(5.0, 2.0)) -- não existe porque ela está privada
}
