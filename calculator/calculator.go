package main

import (
	"fmt"
)

func main() {
	sm := sum(1, 2, 5, 10)
	sb := sub(100, 5)
	dv := div(50, 2)
	mp := mult(10, 10)

	fmt.Println("Soma:", sm)
	fmt.Println("Subtração:", sb)
	fmt.Println("Divisão:", dv)
	fmt.Println("Multiplicação:", mp)
}

func sum(i ...int) int {
	result := 0
	for _, v := range i {
		result += v
	}
	return result
}

func sub(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	result := i[0]
	for _, v := range i[1:] {
		result -= v
	}
	return result
}

func div(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	result := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			fmt.Println("Erro: divisão por zero.")
			return 0
		}
		result /= v
	}
	return result
}

func mult(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	result := 1
	for _, v := range i {
		result *= v
	}
	return result
}
