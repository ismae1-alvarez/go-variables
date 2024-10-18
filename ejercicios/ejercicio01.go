package ejercicios

import "strconv"

func EjerciciosDevolverValores(letras string) (int, string) {

	i, err := strconv.Atoi(letras)

	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}

	if i > 100 {
		return i, "Es mayor a 100"
	} else {
		return i, "Es menor a 100"
	}

}
