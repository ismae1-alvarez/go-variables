package main

import (
	"fmt"
	"udemy-1/ejercicios"
)

func main() {

	// estado, texto := variables.ConviertoaTexto()(121)

	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Si es linux o os", os)
	// } else {
	// 	fmt.Println("Esto es windows", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es linux")
	// case "dasrwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	numero, mensaje := ejercicios.EjerciciosDevolverValores("200")

	fmt.Println(numero)
	fmt.Println(mensaje)

}
