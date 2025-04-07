package pkg

import "math"

type Time []float64
type Voltage []float64

func U(R float64, C float64, Rest float64, I float64, time Time) Voltage {
	tau := R * C
	volt_arr := Voltage{}
	for t := range time {
		sum := Rest + R*I*(1-math.Exp(-time[t]/tau))
		volt_arr = append(volt_arr, sum)
	}
	return volt_arr
}
