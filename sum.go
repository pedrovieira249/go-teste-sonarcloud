package main

import "fmt"

func main() {
	fmt.Println(soma(32, 32))
	fmt.Println(sub(512, 256))
	fmt.Println(multi(32, 16))
	fmt.Println(divisao(1024, 256))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

func divisao(a int, b int) int {
	return a / b
}
