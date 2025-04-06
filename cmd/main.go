package main

import "neuro/iaf-model/internal"
import "fmt"

// Why can't I populate the array outside of the main function here?
var rest [2]float64
var cap [2]float64
var res [2]float64
var x [2]float64
var y [2]float64

func main() {
	rest[0] = -65.0
	rest[1] = -50.0
	cap[0] = 10.0
	cap[1] = 12.0
	res[0] = 5
	res[1] = 6
	x[0] = 1
	y[0] = 1
	x[1] = 2
	y[1] = 2
	// Definitely need a quicker way to create sets of randomized neurons

	a := internal.NewNeuron(rest[0], cap[0], res[0], x[0], y[0])
	b := internal.NewNeuron(rest[1], cap[1], res[1], x[1], y[1])
	distance := a.Distance(b)
	// Square root of 2!
	fmt.Println(distance)
	fmt.Println(a)
}
