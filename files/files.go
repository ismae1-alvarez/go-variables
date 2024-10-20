package files

import (
	"bufio"
	"fmt"
	"os"
	"udemy-1/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.PedirNumero()

	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear e archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.PedirNumero()

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(file string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	arch.Close()

	return true

}

func LeoArchivo() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Hubo un errror leyendo el archivo" + err.Error())
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()

		fmt.Print("> " + registro)
	}

	archivo.Close()
}
