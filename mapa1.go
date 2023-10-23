package main

import (
	"fmt"
	"strings"
)

func contarPalavras(frase string) map[string]int {
	ocorrencias := make(map[string]int)
	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias

}
func main() {
	frase := "A roupa do rato é diferente da roupa do rei do rato"

	fmt.Println(contarPalavras(frase))
}
