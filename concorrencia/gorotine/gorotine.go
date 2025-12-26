package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//a palavra go executa funções de forma assíncrona
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)
	// go fale("Maria", "Ei...", 500) //só vai executar se a thead principal
	// do seu programa estiver funcionando
	// go fale("João", "Opa...", 500)
	// time.Sleep(time.Second * 5)

	// fmt.Println("Fim!")

	go fale("Maria", "Entendi!", 10)
	fale("João", "Parabéns!", 5)
}
