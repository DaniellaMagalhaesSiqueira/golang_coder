package main

import "fmt"

func main() {
	i := 1
	//Go não tem aritmética de ponteiros

	var p *int = nil
	p = &i // pegando o endereço de i passando para o ponteiro

	*p++ //desreferenciado para pegar o valor
	i++

	fmt.Println(p, *p, i)     //0xc00000a0b8 3 3
	fmt.Println(p, *p, i, &i) //0xc00000a0b8 3 3 0xc00000a0b8

}
