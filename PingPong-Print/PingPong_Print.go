// Codigo realizado para Desafio DIO - Go Developers
// Author - lerocha1 (Leandro Rocha)
// Todos direitos Reservados 😊
package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pingar(c)
	go pong(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
