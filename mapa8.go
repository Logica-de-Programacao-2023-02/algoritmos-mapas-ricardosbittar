package main

func calculardespesas(despesas map[string]float64) map[string]float64 {
	totalDespesas := 0.0
	resultado := make(map[string]float64)

	for _, valor := range despesas {
		totalDespesas += valor
	}
	media := totalDespesas / float64(len(despesas))
	for nome, valor := range despesas {
		resultado[nome] = valor - media
	}
	return resultado
}

func gastos(conta map[string]float64) map[string]float64 {

}
