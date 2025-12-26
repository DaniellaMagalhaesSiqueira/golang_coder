package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //inviando o valor um para o canal (escrita)
	// <-ch saindo do canal um valor
	ch <- 2
	fmt.Println(<-ch)
}
