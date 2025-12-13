package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

//interfaces são implementadas implicitamente

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// var carro2 esportivo = ferrari{"F40", false, 0}
	// quando está trabalhando com tipo não tem problema
	// quando está trabalhando com interface precisa ter a referência de memória
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()
	fmt.Println(carro1, carro2)

}
