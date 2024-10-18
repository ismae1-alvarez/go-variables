package main

import (
	"fmt"
	"udemy-1/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(121)

	fmt.Println(estado)
	fmt.Println(texto)

}
