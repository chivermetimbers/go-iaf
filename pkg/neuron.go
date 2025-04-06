package pkg

import "math"
import "gonum.org/v1/gonum/graph"

// TODO: Write unit tests
func (p Neuron) Distance(q Neuron) float64 {
	a := q.X - p.X
	b := q.Y - p.Y
	return math.Sqrt(a*a + b*b)
}

func (n Neuron) NewGraphNode(id int64, name string) *GraphNode {
	return &GraphNode{name: name, id: id, neuron: n}
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

type GraphNode struct {
	id         int64
	neighbours []graph.Node
	roots      []*GraphNode
	name       string
	neuron     Neuron
}

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
