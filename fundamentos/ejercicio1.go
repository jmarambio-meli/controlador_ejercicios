package fundamentos

import "fmt"

func Fundamentos_Ejercicio1() {
	var nombre string
	var direccion string

	fmt.Println("Ingrese su nombre")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su direccion")
	fmt.Scanln(&direccion)

	fmt.Println(nombre, direccion)

}
