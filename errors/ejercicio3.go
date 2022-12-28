package errors

import (
	"errors"
	"fmt"
)

var (
	Error = errors.New("Error: El salario es menor a 10000")
)

func Error_ejercicio3() {
	var salary int = 15000
	err := validarSalario3(salary)
	if err != nil {
		if errors.Is(err, Error) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Debe pagar Impuesto")
	}
	fmt.Println(errors.Is(err, Error))

}

func validarSalario3(salary int) error {
	if salary < 10000 {
		return Error
	}

	return nil
}
