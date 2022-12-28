package estructuras_control

import "fmt"

func Estructuras_Control_Ejercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de benjamin es: %v\n", employees["Benjamin"])

	var contador int = 0
	for _, v := range employees {
		if v > 21 {
			contador++
		}
	}

	fmt.Printf("Existen %v personas mayores de 21\n", contador)
	employees["Federico"] = 25

	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}
