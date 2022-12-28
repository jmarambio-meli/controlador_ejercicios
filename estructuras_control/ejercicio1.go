package estructuras_control

import (
	"fmt"
	//"strings"
)

func Estructuras_Control_Ejercicio1() {
	var palabra string
	fmt.Scanln(&palabra)

	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}

	fmt.Println("El largo de la palabra es: ", len(palabra))
}
