package ejer_interfaces

import (
	"fmt"
	"udemy-1/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	// hu.EstaVivo()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
