package main

import "math"

//inicializando um método ou variável com a primeira letra maíuscula é PÚBLICA
//Inicializando com letra minúscula é PRIVADO //será visível dentro do pacote
// Ponto - gera algo público
//ponto ou _Ponto - gerará algo privado

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Calcula a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
