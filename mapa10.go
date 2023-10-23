package main

import "fmt"

func freqDePares(lista []int) map[int]int {
	ocorrencias := make(map[int]int)

	for _, numero := range lista {
		if numero%2 == 0 {
			ocorrencias[numero]++
		}
	}
	return ocorrencias
}
func main() {
	lista := []int{1, 2, 4, 7, 33, 64, 81, 76, 52, 76, 2, 102, 102, 102}
	fmt.Println(freqDePares(lista))

}

