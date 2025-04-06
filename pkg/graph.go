package pkg

import "gonum.org/v1/gonum/graph"

type GraphNode struct {
	id         int64
	neighbours []graph.Node
	roots      []*GraphNode
	// Probably don't need names
	name   string
	neuron Neuron
}

func (n Neuron) NewGraphNode(id int64, name string) *GraphNode {
	return &GraphNode{name: name, id: id, neuron: n}
}
