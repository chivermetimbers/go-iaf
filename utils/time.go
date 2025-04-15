package utils

import (
	"errors"
)

func CreateTimeArray(min_t float64, max_t float64, step float64) ([]float64, error) {
	if step < 0 {
		return nil, errors.New("Negative time step")
	}
	time_points := []float64{}
	current := min_t
	time_points = append(time_points, current)
	for {
		current += step
		if (current + step) > max_t {
			break
		}
		time_points = append(time_points, current)
	}
	return time_points, nil
}
