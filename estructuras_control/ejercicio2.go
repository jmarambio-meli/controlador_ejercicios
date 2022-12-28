package estructuras_control

import (
	"fmt"
)

func Estructuras_Control_Ejercicio2() {

	var edad int
	var respuesta string
	var salario int

	fmt.Println("Cual es su edad?")
	fmt.Scanln(&edad)
	if edad < 22 {
		fmt.Println("No se puede dar prestamos a personas menor de 22")
		return
	}
	fmt.Println("Se encuentra empleado? Si es afirmativo presione 'S', sino 'N' ")
	fmt.Scanln(&respuesta)
	if respuesta == "N" {
		fmt.Println("No se puede dar prestamos a personas desempleadas")
		return
	}
	fmt.Println("Lleva mas de un año trabajando? Si es afirmativo presione 'S', sino 'N' ")
	fmt.Scanln(&respuesta)
	if respuesta == "N" {
		fmt.Println("No se puede dar prestamos a personas que llevan menos de 1 año de antiguedad laboral")
		return
	}
	fmt.Println("Cuanto es su salario?")
	fmt.Scanln(&salario)
	if salario < 100000 {
		fmt.Println("Usted puede recibir un prestamo pero con intereses")
	} else {
		fmt.Println("Usted puede recibir un prestamo sin intereses")
	}
}
