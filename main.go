package main

import (
	"fmt"
	"primeiroProjeto/calc"
)

func main() {
	num1 := 4
	num2 := 0

	resultado, err := calc.Calculadora(num1, num2, "-")

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(resultado)
}
