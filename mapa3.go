package main

import "fmt"

func somarValores(x map[string]int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}
	return soma
}
func main() {
	numeros := map[string]int{
		"1": 2,
		"2": 8,
		"3": 10,
	}
	fmt.Println(somarValores(numeros))
}
