package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf((32000)))

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
}
