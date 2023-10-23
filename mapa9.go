package main

func Fibonacci(n int) map[int]int {
	sequencia := make(map[int]int)
	a, b := 0, 1

	for i := 0; a <= n; i++ {
		sequencia[i] = a
		a, b = b, a+b
	}
	return sequencia
}