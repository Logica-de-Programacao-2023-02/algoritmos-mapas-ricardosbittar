package main

import "strings"

func contarLetras(palavra string) map[string]int {
	ocorrencias := make(map[string]int)
	for _, letra := range palavra {
		letras := string(letra)
		ocorrencias[letras]++
	}
	return ocorrencias
}
func dividirPalavras(frase string) map[string]map[string]int {
	resultado := make(map[string]map[string]int)
	palavras := strings.Fields(frase)
	for _, palavra := range palavras {
		resultado[palavra] = contarLetras(palavra)
	}
	return resultado

}
