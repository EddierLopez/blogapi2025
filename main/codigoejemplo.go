package main

import (
	"fmt"
)

func menu() {
	fmt.Println("----Menu de opciones----")
	fmt.Println("1- Sumar")
	fmt.Println("2- Obtener mayor y menor")
	fmt.Println("3- Salir del sistema")
}

/**Funcion sumar*/
func sumar(num1, num2 int) int {
	var result int //resultado
	result = num1 + num2
	return result
}

func mayorMenor(maxEntradas int) (mayor int, menor int) {
	var entrada int
	for i := 1; i <= maxEntradas; i++ {
		fmt.Println("Entrada #", i)
		fmt.Scan(&entrada)
		if i == 1 {
			mayor = entrada
			menor = entrada
		} else {
			switch {
			case entrada > mayor:
				mayor = entrada
			case entrada < menor:
				menor = entrada
			}
		}
	}

	return
}

func main() {
	var opt uint8
	var resultado, n1, n2 int
	var mayor, menor, cant int
	for {
		menu()
		fmt.Println("Digite una opción:")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			fmt.Println("Digite el primer numero:")
			fmt.Scan(&n1)
			fmt.Println("Digite el segundo numero:")
			fmt.Scan(&n2)
			resultado = sumar(n1, n2)
			fmt.Println("El resultado es:", resultado)
		case 2:
			fmt.Println("Digite la cantidad de numeros a ingresar")
			fmt.Scan(&cant)
			mayor, menor = mayorMenor(cant)
			fmt.Println("El mayor es:", mayor, " y el menor es:", menor)

		}
		if opt == 3 {
			break
		}

	}
}
