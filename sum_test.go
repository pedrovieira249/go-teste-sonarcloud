package main

// Para rodar os testes, use: go test -coverprofile=coverage.out
// Para ver o que está faltando, use: go tool cover -func=coverage.out
import (
	"testing"
)

func TestSoma(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"soma positivos", 15, 15, 30},
		{"soma com zero", 10, 0, 10},
		{"soma negativos", -5, -3, -8},
		{"soma positivo e negativo", 10, -5, 5},
		{"soma zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := soma(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("soma(%d, %d) = %d; esperado %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"subtração simples", 10, 5, 5},
		{"subtração com resultado negativo", 5, 10, -5},
		{"subtração com zero", 10, 0, 10},
		{"subtração de negativos", -5, -3, -2},
		{"subtração zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sub(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("sub(%d, %d) = %d; esperado %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMulti(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"multiplicação simples", 5, 3, 15},
		{"multiplicação com zero", 10, 0, 0},
		{"multiplicação com negativos", -5, 3, -15},
		{"multiplicação de negativos", -4, -2, 8},
		{"multiplicação por um", 7, 1, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := multi(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("multi(%d, %d) = %d; esperado %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// validateDivisaoError verifica se há erro quando esperado
func validateDivisaoError(t *testing.T, a, b int, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("divisao(%d, %d) deveria retornar erro, mas não retornou", a, b)
	}
}

// validateDivisaoSuccess verifica resultado quando não há erro
func validateDivisaoSuccess(t *testing.T, a, b, result, expected int, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("divisao(%d, %d) retornou erro inesperado: %v", a, b, err)
		return
	}
	if result != expected {
		t.Errorf("divisao(%d, %d) = %d; esperado %d", a, b, result, expected)
	}
}

func TestDivisao(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectError bool
	}{
		{"divisão simples", 10, 2, 5, false},
		{"divisão exata", 100, 10, 10, false},
		{"divisão com resto (truncada)", 10, 3, 3, false},
		{"divisão por um", 7, 1, 7, false},
		{"divisão de zero", 0, 5, 0, false},
		{"divisão por zero - deve dar erro", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divisao(tt.a, tt.b)

			if tt.expectError {
				validateDivisaoError(t, tt.a, tt.b, err)
			} else {
				validateDivisaoSuccess(t, tt.a, tt.b, result, tt.expected, err)
			}
		})
	}
}

// TestDivisaoComErro testa explicitamente o branch de erro na divisão
func TestDivisaoComErro(t *testing.T) {
	_, err := divisao(10, 0)
	if err == nil {
		t.Error("divisao(10, 0) deveria retornar erro")
	}
	if err.Error() != "divisão por zero não é permitida" {
		t.Errorf("mensagem de erro incorreta: %v", err)
	}
}

// TestDivisaoSemErro testa explicitamente o branch de sucesso na divisão
func TestDivisaoSemErro(t *testing.T) {
	result, err := divisao(10, 2)
	if err != nil {
		t.Errorf("divisao(10, 2) não deveria retornar erro: %v", err)
	}
	if result != 5 {
		t.Errorf("divisao(10, 2) = %d; esperado 5", result)
	}
}

// TestRunCalculations testa a função runCalculations
func TestRunCalculations(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("runCalculations() causou panic: %v", r)
		}
	}()

	// Executa runCalculations() que inclui:
	// - soma, sub, multi (branch de sucesso)
	// - divisao com sucesso (branch else)
	// - divisao com erro (branch if err != nil)
	runCalculations()
}

// TestMain testa a função main diretamente
func TestMain(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() causou panic: %v", r)
		}
	}()

	main()
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(10, 20)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub(20, 10)
	}
}

func BenchmarkMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multi(10, 20)
	}
}

func BenchmarkDivisao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divisao(100, 10)
	}
}

func BenchmarkDivisaoComErro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divisao(100, 0)
	}
}
