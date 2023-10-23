package main

import "fmt"

func juntarMapas(x, y map[string]int) map[string]int {
	resultado := make(map[string]int)
	for chave1, valor1 := range x {
		resultado[chave1] = valor1
	}
	for chave2, valor2 := range y {
		for chave3 := range resultado {
			if y[chave2] != resultado[chave3] {
				resultado[chave2] = valor2
			}
			if y[chave2] == resultado[chave3] {

			}
			resultado[chave2] = valor2

		}

	}
	return resultado

}
func main() {
	x := map[string]int{
		"camas":    3,
		"cadeiras": 4,
		"mesas":    1,
	}

	y := map[string]int{
		"camas":    4,
		"armarios": 1,
	}
	fmt.Println(juntarMapas(x, y))
}
