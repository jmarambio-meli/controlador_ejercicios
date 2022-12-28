package estructuras_metodos_composicion

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	IdEmployee int
	Position   string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Num Trabajador: %v, Posición: %s\nId Persona: %v, Nombre: %s, Fecha de nacimiento: %s\n", e.IdEmployee, e.Position, e.Id, e.Name, e.DateOfBirth)
}

func Estructuras_Metodos_Composicion_Ejercicio2() {
	person1 := Person{
		Id:          20777519,
		Name:        "Jean",
		DateOfBirth: "30/05/2021",
	}

	employee1 := Employee{
		IdEmployee: 1,
		Position:   "Software Developer",
		Person:     person1,
	}

	employee1.PrintEmployee()
}
