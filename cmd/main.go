package main

import (
	"fmt"
	"log"
	"neuro/iaf-model/pkg"
	"neuro/iaf-model/utils"

	"gonum.org/v1/plot"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//import "gonum.org/v1/plot"

//TODO: Add AddToGraph function

// Why can't I populate the array outside of the main function here?
var rest [2]float64
var cap [2]float64
var res [2]float64
var x [2]float64
var y [2]float64
var i [2]float64

func main() {
	rest[0] = -20.0
	rest[1] = -50.0
	cap[0] = 10.0
	cap[1] = 12.0
	res[0] = 5
	res[1] = 6
	x[0] = 1
	y[0] = 1
	x[1] = 2
	y[1] = 2
	i[0] = 10
	i[1] = 1
	// Definitely need a quicker way to create sets of randomized neurons
	time_points := utils.CreateTimeArray(0, 30, 0.25)
	a := pkg.NewNeuron(rest[0], cap[0], res[0], x[0], y[0], i[0])
	//b := pkg.NewNeuron(rest[1], cap[1], res[1], x[1], y[1], u[1])

	membrane_eq := pkg.U(a.R, a.C, a.Rest, a.I, time_points)
	fmt.Printf("There are %v time points", len(time_points))
	fmt.Printf("There are %v solutions", len(membrane_eq))
	p := plot.New()

	pts := make(plotter.XYs, len(time_points))
	for i := range time_points {
		pts[i].X = time_points[i]
		pts[i].Y = membrane_eq[i]

	}

	//Add scatter plot
	points, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(points)

	if err := p.Save(10.16*vg.Centimeter, 10.16*vg.Centimeter, "plot.png"); err != nil {
		log.Fatal(err)
	}

}
