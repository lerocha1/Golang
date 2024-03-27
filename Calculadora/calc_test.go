package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	a := 10.0
	b := 5.0
	expected := 15.0
	result := soma(a, b)
	if result != expected {
		t.Errorf("Soma incorreta. Esperado: %.2f, Resultado: %.2f", expected, result)
	}
}

func TestSubtrai(t *testing.T) {
	a := 10.0
	b := 5.0
	expected := 5.0
	result := subtrai(a, b)
	if result != expected {
		t.Errorf("Subtração incorreta. Esperado: %.2f, Resultado: %.2f", expected, result)
	}
}

func TestMultiplica(t *testing.T) {
	a := 10.0
	b := 5.0
	expected := 50.0
	result := multiplica(a, b)
	if result != expected {
		t.Errorf("Multiplicação incorreta. Esperado: %.2f, Resultado: %.2f", expected, result)
	}
}

func TestDivisao(t *testing.T) {
	type test struct {
		a         float64
		b         float64
		expected  float64
		shouldErr bool
	}

	tests := []test{
		{10.0, 5.0, 2.0, false},
		{10.0, 0.0, 0.0, true},
	}

	for _, tc := range tests {
		result, err := divisao(tc.a, tc.b)
		if tc.shouldErr {
			if err == nil {
				t.Errorf("Esperava um erro ao dividir %.2f por %.2f, mas não ocorreu", tc.a, tc.b)
			}
		} else {
			if err != nil {
				t.Errorf("Erro inesperado ao dividir %.2f por %.2f: %v", tc.a, tc.b, err)
			}
			if result != tc.expected {
				t.Errorf("Divisão incorreta. Esperado: %.2f, Resultado: %.2f", tc.expected, result)
			}
		}
	}
}
