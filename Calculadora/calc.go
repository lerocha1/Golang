package main

import (
	"errors"
	"fmt"
)

// Add retorna a soma de dois números.
func soma(a, b float64) float64 {
	return a + b
}

// Subtract retorna a subtração de dois números.
func subtrai(a, b float64) float64 {
	return a - b
}

// Multiply retorna a multiplicação de dois números.
func multiplica(a, b float64) float64 {
	return a * b
}

// Divide retorna a divisão de dois números.
// Retorna um erro se a divisão por zero for tentada.
func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não pode dividir por zero")
	}
	return a / b, nil
}

func main() {
	a := 10.0
	b := 5.0
	fmt.Printf("%.2f\n", soma(a, b))
	fmt.Printf("%.2f\n", subtrai(a, b))
	fmt.Printf("%.2f\n", multiplica(a, b))
	result, _ := divisao(a, b)
	fmt.Printf("%.2f\n", result)
}
