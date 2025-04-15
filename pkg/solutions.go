package pkg

import "math"

// This is the solution for constant I
func SolutionConstI(R float64, C float64, Rest float64, I float64, time []float64) []float64 {
	// Apparently negative R & C are allowed, and they have implications.
	// I wonder if the bound for the values of R in a real neuron includes negative values
	tau := R * C
	volt_arr := []float64{}
	volt_arr = append(volt_arr, Rest)
	for i := range time {
		sum := Rest + R*I*(1-math.Exp(-time[i]/tau))
		volt_arr = append(volt_arr, sum)
	}
	return volt_arr
}
