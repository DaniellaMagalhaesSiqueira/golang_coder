package main

import (
	"fmt"
)

func rotina(ch chan int) {
	// fmt.Println("Executou!")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou depois do quatro")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) //buffer de 3 posições
	go rotina(ch)

	// time.Sleep(time.Second)
	fmt.Println(<-ch) //libera um espaço no buffer

}
