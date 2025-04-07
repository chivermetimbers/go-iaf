package utils

import "neuro/iaf-model/pkg"

func CreateTimeArray(min_t float64, max_t float64, step float64) pkg.Time {
	time_points := pkg.Time{}
	current := min_t
	time_points = append(time_points, min_t)
	for {
		current += step
		if (current + step) > max_t {
			break
		}
		time_points = append(time_points, current)
	}
	return time_points

}
