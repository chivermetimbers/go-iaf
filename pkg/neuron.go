package pkg

import "math"

// TODO: Write unit tests
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
	U    float64
}

// should probably separate this out into a separate file at some point

// Turn this into a method
func NewNeuron(resting_potential float64, capacitance float64, resistance float64, x float64, y float64, u float64) Neuron {
	return Neuron{
		Rest: resting_potential,
		C:    capacitance,
		R:    resistance,
		X:    x,
		Y:    y,
		U:    u,
	}
}

// This will eventually compute the potential for a given time
// Or maybe it would be better if took a time array and returned a voltage array?
func (n Neuron) NullPotential(t float64) float64 {
	return t
}
