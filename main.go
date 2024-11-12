package main

import (
	"fmt"
	"ramiroguzmanc/go-intermedio/employee"
)

func main() {
	e := employee.Employee{}
	e.SetId(1)
	e.SetName("Sergio")

	fmt.Println(e.GetName())

	e2 := new(employee.Employee) // Cuando se utiliza new, regresa un apuntador
	fmt.Println(*e2)

	e3 := employee.NewEmployee(3, "Rama", true)
	fmt.Println(*e3)

}
