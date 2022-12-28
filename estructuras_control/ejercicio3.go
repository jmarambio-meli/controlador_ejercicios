package estructuras_control

import (
	"fmt"
)

func Estructuras_Control_Ejercicio3() {
	meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	var numeroMes int
	fmt.Println("Ingrese el numero de mes")
	fmt.Scanln(&numeroMes)
	if numeroMes > 1 && numeroMes < 12 {
		fmt.Printf("El mes con el numero %v es: %v\n", int(numeroMes), meses[numeroMes-1])
	} else {
		fmt.Printf("el numero ingresado es incorrecto\n")
	}
}
