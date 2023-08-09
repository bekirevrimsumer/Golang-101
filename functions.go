package main

import "fmt"

func main() {
	x := []float64{ 98, 93, 77, 82, 83 }
	fmt.Println("average:", average(x))
}

func average(x []float64) float64 {
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	return total / float64(len(x))
}