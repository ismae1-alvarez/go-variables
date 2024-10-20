package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 3, 4}
var arreglo [10]int = [10]int{6, 4, 45, 53, 34, 35, 43, 45}

func MuestroSlices() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  // Slices creado desde un vector, dedes ña posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // Slices creado desde un vector,  desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, capcidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, capcidad %d", len(nums), cap(nums))

}
