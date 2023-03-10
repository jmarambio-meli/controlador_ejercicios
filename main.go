package main

import (
	"fmt"
	"github/jmarambio-meli/errors"
	"github/jmarambio-meli/estructuras_control"
	"github/jmarambio-meli/estructuras_metodos_composicion"
	"github/jmarambio-meli/funciones"
	"github/jmarambio-meli/fundamentos"
	"github/jmarambio-meli/panics_defer_recover"
	"github/jmarambio-meli/punteros"
)

func main() {
	var opcion string
	loop := true
	for loop {
		println("Ingrese 1 Para acceder a los ejercicios del Día 1")
		println("Ingrese 2 Para acceder a los ejercicios del Día 2")
		println("Ingrese 3 Para acceder a los ejercicios del Día 3")
		println("Ingrese 4 Para acceder a los ejercicios del Día 4")
		println("Ingrese S Para salir del Menu")
		fmt.Printf("\n")
		fmt.Scanln(&opcion)

		switch opcion {
		case "1":
			fmt.Println("Ingrese 1 Para acceder al ejercicio n°1 de fundamentos")
			fmt.Println("Ingrese 2 Para acceder al ejercicio n°2 de fundamentos")

			fmt.Println("Ingrese 3 Para acceder al ejercicio n°1 de estructuras_basicas")
			fmt.Println("Ingrese 4 Para acceder al ejercicio n°2 de estructuras_basicas")
			fmt.Println("Ingrese 5 Para acceder al ejercicio n°3 de estructuras_basicas")
			fmt.Println("Ingrese 6 Para acceder al ejercicio n°4 de estructuras_basicas")
			fmt.Scanln(&opcion)
			switch opcion {
			case "1":
				fundamentos.Fundamentos_Ejercicio1()
			case "2":
				fundamentos.Fundamentos_Ejercicio2()
			case "3":
				estructuras_control.Estructuras_Control_Ejercicio1()
			case "4":
				estructuras_control.Estructuras_Control_Ejercicio2()
			case "5":
				estructuras_control.Estructuras_Control_Ejercicio3()
			case "6":
				estructuras_control.Estructuras_Control_Ejercicio4()
			}
		case "2":
			fmt.Println("Ingrese 1 Para acceder al ejercicio n°1 de funciones")
			fmt.Println("Ingrese 2 Para acceder al ejercicio n°2 de funciones")
			fmt.Println("Ingrese 3 Para acceder al ejercicio n°3 de funciones")
			fmt.Println("Ingrese 4 Para acceder al ejercicio n°4 de funciones")
			fmt.Println("Ingrese 5 Para acceder al ejercicio n°5 de funciones")
			fmt.Scanln(&opcion)
			switch opcion {
			case "1":
				funciones.Funciones_Ejercicio1()
			case "2":
				funciones.Funciones_Ejercicio2()
			case "3":
				funciones.Funciones_Ejercicio3()
			case "4":
				funciones.Funciones_Ejercicio4()
			case "5":
				funciones.Funciones_Ejercicio5()
			}
		case "3":
			fmt.Println("Ingrese 1 Para acceder al ejercicio n°1 de estructuras, metodos y composicion")
			fmt.Println("Ingrese 2 Para acceder al ejercicio n°2 de estructuras, metodos y composicion")
			fmt.Println("Ingrese 3 Para acceder al ejercicio n°1 de punteros")
			fmt.Println("Ingrese 4 Para acceder al ejercicio n°2 de punteros")
			fmt.Scanln(&opcion)
			switch opcion {
			case "1":
				estructuras_metodos_composicion.Estructuras_Metodos_Composicion_Ejercicio1()
			case "2":
				estructuras_metodos_composicion.Estructuras_Metodos_Composicion_Ejercicio2()
			case "3":
				punteros.Punteros_Ejercicio1()
			case "4":
				punteros.Punteros_Ejercicio2()
			}
		case "4":
			fmt.Println("Ingrese 1 Para acceder al ejercicio n°1 de errors")
			fmt.Println("Ingrese 2 Para acceder al ejercicio n°2 de errors")
			fmt.Println("Ingrese 3 Para acceder al ejercicio n°3 de errors")
			fmt.Println("Ingrese 4 Para acceder al ejercicio n°4 de errors")
			fmt.Println("Ingrese 5 Para acceder al ejercicio n°5 de errors")
			fmt.Println("Ingrese 6 Para acceder al ejercicio n°1 de panics_defer_recover")
			fmt.Println("Ingrese 7 Para acceder al ejercicio n°2 de panics_defer_recover")
			fmt.Println("Ingrese 8 Para acceder al ejercicio n°3 de panics_defer_recover")
			fmt.Scanln(&opcion)
			switch opcion {
			case "1":
				errors.Error_ejercicio1()
			case "2":
				errors.Error_ejercicio2()
			case "3":
				errors.Error_ejercicio3()
			case "4":
				errors.Error_ejercicio4()
			case "5":
				errors.Error_ejercicio5()
			case "6":
				panics_defer_recover.Panics_defer_recover_ejercicio1()
			case "7":
				panics_defer_recover.Panics_defer_recover_ejercicio2()
			case "8":
				panics_defer_recover.Panics_defer_recover_ejercicio3()
			}
		}

	}
}
