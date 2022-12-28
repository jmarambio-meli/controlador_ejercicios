package panics_defer_recover

import (
	"fmt"
	"os"
)

func Panics_defer_recover_ejercicio1() {

	defer func() {
		fmt.Println("ejecución finalizada")
	}()

	_, err := os.Open("customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println("fin")
}
