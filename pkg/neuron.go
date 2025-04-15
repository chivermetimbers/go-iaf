package pkg

type Neuron struct {
	Rest float64
	C    float64
	R    float64
	I    float64
}

func NewNeuron(resting_potential float64, capacitance float64, resistance float64, i float64) Neuron {
	return Neuron{
		Rest: resting_potential,
		C:    capacitance,
		R:    resistance,
		I:    i,
	}
}
