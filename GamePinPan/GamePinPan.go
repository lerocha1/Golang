//Codigo realizado para Desafio DIO - Go Developers
//Author - lerocha1 (Leandro Rocha)
//Todos direitos Reservados 😊

//Desafio Pin e Pan - Game
//Numero de 1 a 100, para multiplos de 3 imprimi Pin e para multiplos de 5 imprimi Pan e se for multiplo de 3 e 5 imprimi PinPan

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			continue
		}
	}
}
