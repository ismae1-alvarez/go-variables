package main

import (
	"udemy-1/middleware"
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

	// numero, mensaje := ejercicios.EjerciciosDevolverValores("200")

	// fmt.Println(numero)
	// fmt.Println(mensaje)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.PedirNumero())

	// files.SumaTabla()
	// files.LeoArchivo()

	// files.LeoArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()

	// funciones.Exponencia(2)

	// arreglos_slices.MuestraArreglos()
	// arreglos_slices.MuestroSlices()
	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()

	// users.AltaUsuario()

	// Pedro := new(modelos.Hombre)

	// ejer_interfaces.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Maria)

	// defer_panic.EjemploPanic()

	// canal1 := make(chan bool)

	// go goroutines.MiNombreLento("ismael", canal1)

	// defer func() {
	// 	<-canal1
	// }()

	// fmt.Println("Estoy aqui")

	// webserver.MiWebServer()
	middleware.MiMiddleware()

}
