//Codigo realizado para Desafio DIO - Go Developers
//Author - lerocha1 (Leandro Rocha)
//Todos direitos Reservados 😊

package main

import (
	"fmt"
)

// Type para armazenar a lista de liquidos e temperaturas de Ebolição em Kelvin
type Liquid struct {
	Nome   string
	Kelvin float64
}

func main() {
	// Defina os líquidos e suas temperaturas de ebulição em Kelvin. Listei apenas 10 para exemplificar.
	liquido := []Liquid{
		{Nome: "agua", Kelvin: 373},
		{Nome: "mercurio", Kelvin: 629.87},
		{Nome: "nitrogenio", Kelvin: 77.35},
		{Nome: "helio", Kelvin: 4.22},
		{Nome: "fluor", Kelvin: 85.03},
		{Nome: "argonio", Kelvin: 87.30},
		{Nome: "gálio", Kelvin: 2403.0},
		{Nome: "iodo", Kelvin: 457.40},
		{Nome: "enxofre", Kelvin: 717.87},
		{Nome: "chumbo", Kelvin: 2022.0},
	}

	// Com uma iteração for, pergunte ao usuario qual liquido ele deseja saber o ponto de Ebulição.
	//Aqui listamos na tela.
	fmt.Println("Escolha um número equivalente a um dos seguintes líquidos:")
	for i, l := range liquido {
		fmt.Printf("%d. %s\n", i+1, l.Nome)
	}

	// Armazena o liquido escolhido pelo usuario, dentros os 10 disponiveis.
	var escolha int
	fmt.Print("Digite o número da opção desejada (apenas  numero): ")
	fmt.Scan(&escolha)
	fmt.Println("********************************************************************************")

	// Verifica se nao digitou uma opção diferente da fornecida, para garantir que nao haja erros
	if escolha < 1 || escolha > len(liquido) {
		fmt.Println("Opção inválida. Tente novamente.")
		return
	}

	// Por ultimo, calcule e imprime a temperatura em Celsius C = k - 273
	celsius := liquido[escolha-1].Kelvin - 273
	fmt.Printf("A temperatura de ebulição de %s em Celsius é %.2f°C.\n", liquido[escolha-1].Nome, celsius)
}
