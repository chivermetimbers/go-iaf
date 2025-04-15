package utils

import (
	"log"
)

// This is to be used for non-linear models
// NOTE: Need to ensure function is in the correct form y' = f(x,y)
func RungeKutta(time []float64, mbpot func([]float64) ([]float64, error), strt float64, fin float64, stp float64) ([]float64, error) {
	volt, err := mbpot(time)
	if err != nil {
		log.Fatal(err)
	}
	// Code goes here
	return volt, err
}
