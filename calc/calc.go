package calc

import (
	"errors"
)

func Calculadora(a int, b int, op string) (int, error) {
	if op == "x" {
		return (a * b), nil

	} else if op == "+" {

		return (a + b), nil

	} else if op == "-" {

		return (a - b), nil

	} else if op == ":" {

		if b == 0 {
			return 0, errors.New("O denominador não pode ser zero")
		}

		return (a / b), nil

	} else {

		return 0, errors.New("O Operador passado não existe!!!")
	}
}
