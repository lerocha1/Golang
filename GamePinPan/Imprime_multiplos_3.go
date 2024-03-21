//Codigo realizado para Desafio DIO - Go Developers
//Author - lerocha1 (Leandro Rocha)
//Todos direitos Reservados 😊

//Usar operador % e o laço for!
//Exibir os números de 1 a 100 que são divisíveis por 3.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
