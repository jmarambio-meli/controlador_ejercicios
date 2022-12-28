package estructuras_metodos_composicion

import (
	"fmt"
)

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func Estructuras_Metodos_Composicion_Ejercicio1() {
	product := Product{}

	product1 := Product{1, "manzana", 1000, "descripcion", "frutas"}
	product2 := Product{2, "platano", 500, "descripcion", "frutas"}
	product3 := Product{3, "durazno", 1500, "descripcion", "frutas"}
	product4 := Product{4, "pera", 2000, "descripcion", "frutas"}

	product1.Save()
	product2.Save()
	product3.Save()
	product4.Save()

	product.GetAll()
	product.GetById(3)
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, v := range Products {
		fmt.Printf("Id: %v, Nombre: %s, Price: %f, Description: %s, Category: %s\n\n", v.Id, v.Name, v.Price, v.Description, v.Category)
	}
}

func (p Product) GetById(id int) {
	for _, v := range Products {
		if id == v.Id {
			fmt.Println("!!Producto con encontrado!!")
			fmt.Printf("Id: %v, Nombre: %s, Price: %f, Description: %s, Category: %s\n", v.Id, v.Name, v.Price, v.Description, v.Category)
			break
		}
	}
}
