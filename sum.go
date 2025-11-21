package main

import (
	"errors"
	"fmt"
)

func main() {
	runCalculations()
}

// runCalculations executa todas as operações matemáticas
// Esta função pode ser testada facilmente
func runCalculations() {
	fmt.Println("Soma:", soma(32, 32))
	fmt.Println("Subtração:", sub(512, 256))
	fmt.Println("Multiplicação:", multi(32, 16))

	result, err := divisao(1024, 256)
	if err != nil {
		fmt.Println("Erro na divisão:", err)
	} else {
		fmt.Println("Divisão:", result)
	}

	_, err = divisao(10, 0)
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
}

// soma retorna a soma de dois números inteiros
func soma(a, b int) int {
	return a + b
}

// sub retorna a subtração de dois números inteiros
func sub(a, b int) int {
	return a - b
}

// multi retorna a multiplicação de dois números inteiros
func multi(a, b int) int {
	return a * b
}

// divisao retorna a divisão de dois números inteiros
// Retorna erro se o divisor for zero
func divisao(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}
