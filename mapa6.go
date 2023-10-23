package main

func somarMapas(mapas []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, mapa := range mapas {
		for palavra, contagem := range mapa {
			soma[palavra] += contagem
		}
	}

	return soma
}
