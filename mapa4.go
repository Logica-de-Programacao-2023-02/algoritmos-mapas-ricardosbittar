package main

import "sort"

func anagramas(palavras []string) map[string][]string {

	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		palavraOrdenada := sort.Strings(palavra)

		if _, ok := anagramas[palavraOrdenada]; !ok {

			anagramas[palavraOrdenada] = []string{palavra}
		} else {

			anagramas[palavraOrdenada] = append(anagramas[palavraOrdenada], palavra)
		}
	}

	return anagramas
}
