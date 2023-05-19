package main

import "fmt"

func main() {
	notasAlunos := map[string][]float64{
		"João":  []float64{5, 8, 6, 7},
		"Pedro": []float64{6, 4.5, 5.5, 6},
		"Maria": []float64{7.5, 8, 7.5, 5.5},
	}

	fmt.Println(notas(notasAlunos))
}

func notas(notasAlunos map[string][]float64) map[string]float64 {
	mediaAlunos := make(map[string]float64)

	for a, notas := range notasAlunos {
		var sum float64
		for n := range notas {
			sum += notas[n]
		}
		mediaAlunos[a] = sum / 4
	}

	return mediaAlunos
}
