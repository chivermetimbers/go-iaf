package utils

import "math"

func MinFloat64(slice []float64) float64 {
	if len(slice) == 0 {
		return math.NaN()
	}
	min := slice[0]
	for _, v := range slice[1:] {
		min = math.Min(min, v)
	}
	return min
}

func MaxFloat64(slice []float64) float64 {
	if len(slice) == 0 {
		return math.NaN() // or panic/return error for empty slice
	}
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}
