package neuron

type Neuron struct {
	rest float64
	cap  float64
	res  float64
}

func newNeuron(resting_potential float64, capacitance float64, resistance float64) Neuron {
	return Neuron{
		rest: resting_potential,
		cap:  capacitance,
		res:  resistance,
	}
}
