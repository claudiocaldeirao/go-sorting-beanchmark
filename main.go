package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"benchmark/sort"

	"github.com/wcharczuk/go-chart/v2"
)

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100000)
	}
	return arr
}

func cloneArray(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	return copyArr
}

func main() {
	sizes := []int{100, 1000, 10000}
	algorithms := []string{"Quick Sort", "Merge Sort", "Heap Sort", "Bubble Sort", "Insertion Sort"}

	// Mapa de tempos por algoritmo
	times := map[string][]float64{
		"Quick Sort":     {},
		"Merge Sort":     {},
		"Heap Sort":      {},
		"Bubble Sort":    {},
		"Insertion Sort": {},
	}

	for _, size := range sizes {
		base := generateRandomArray(size)

		for _, algo := range algorithms {
			var arr = cloneArray(base)
			start := time.Now()

			switch algo {
			case "Quick Sort":
				sort.QuickSort(arr)
			case "Merge Sort":
				sort.MergeSort(arr)
			case "Heap Sort":
				sort.HeapSort(arr)
			case "Bubble Sort":
				sort.BubbleSort(arr)
			case "Insertion Sort":
				sort.InsertionSort(arr)
			}

			elapsed := time.Since(start).Seconds()
			times[algo] = append(times[algo], elapsed)
		}
	}

	// Construir gráfico
	var series []chart.Series
	for _, algo := range algorithms {
		x := []float64{}
		y := times[algo]
		for _, size := range sizes {
			x = append(x, float64(size))
		}
		series = append(series, chart.ContinuousSeries{
			Name:    algo,
			XValues: x,
			YValues: y[:len(x)],
		})
	}

	graph := chart.Chart{
		Title:  "Tempo de Execução de Algoritmos de Ordenação",
		XAxis:  chart.XAxis{Name: "Tamanho do Array"},
		YAxis:  chart.YAxis{Name: "Tempo (segundos)"},
		Series: series,
		Elements: []chart.Renderable{
			chart.LegendLeft(&chart.Chart{
				Series: series,
			}),
		},
	}

	filename := "benchmark_chart.png"
	f, _ := os.Create(filename)
	defer f.Close()
	graph.Render(chart.PNG, f)
	fmt.Printf("Gráfico gerado: %s \n", filename)
}
