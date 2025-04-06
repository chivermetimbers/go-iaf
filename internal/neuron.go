package internal

import "math"

func (p Neuron) Distance(q Neuron) float64 {
	a := q.X - p.X
	b := q.Y - p.Y
	return math.Sqrt(a*a + b*b)
}

type Neuron struct {
	Rest float64
	C    float64
	R    float64
	X    float64
	Y    float64
}

// Turn this into a method
func NewNeuron(resting_potential float64, capacitance float64, resistance float64, x float64, y float64) Neuron {
	return Neuron{
		Rest: resting_potential,
		C:    capacitance,
		R:    resistance,
		X:    x,
		Y:    y,
	}
}
