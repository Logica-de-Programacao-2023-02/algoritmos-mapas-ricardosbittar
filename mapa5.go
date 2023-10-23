package main

import (
	"fmt"
	"strings"
)

func contador(frase string) map[string]int {
	palavras := strings.Fields(frase)
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		for _, letra := range palavra {
			letras := string(letra)
			ocorrencias[letras]++
		}
	}

	return ocorrencias
}
func main() {
	frase := "gosto de alface, mas nao gosto de tomate"
	fmt.Println(contador(frase))
}
