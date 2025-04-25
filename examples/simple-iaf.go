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

//TODO: Add AddToGraph function

var rest float64
var cap float64
var res float64
var x float64
var y float64
var i float64

func main() {
	rest = -65.0
	cap = 0.10
	res = 0.50
	i = 0.015

	// Create an array of time points
	time_points, err := utils.CreateTimeArray(0, 1, 0.001)
	if err != nil {
		log.Fatal(err)
	}
	a := pkg.NewNeuron(rest, cap, res, i)

	membrane_eq := pkg.SolutionConstI(a.R, a.C, a.Rest, a.I, time_points)
	fmt.Printf("There are %v time points", len(time_points))
	fmt.Printf("There are %v solutions", len(membrane_eq))

	p := plot.New()
	p.Title.Text = "Membrane Voltage (V) vs Time (s)"
	p.X.Label.Text = "Time(s)"
	p.Y.Label.Text = "Membrane Voltage(V)"
	minY := utils.MinFloat64(membrane_eq)
	maxY := utils.MaxFloat64(membrane_eq)
	p.Y.Tick.Marker = plot.ConstantTicks{
		{Value: minY, Label: fmt.Sprintf("%.2f", minY)},
		{Value: (minY + maxY) / 2, Label: fmt.Sprintf("%.2f", (minY+maxY)/2)},
		{Value: maxY, Label: fmt.Sprintf("%.2f", maxY)},
	}

	pts := make(plotter.XYs, len(time_points))
	for i := range time_points {
		pts[i].X = time_points[i]
		pts[i].Y = membrane_eq[i]

	}

	//Add Line plot
	points, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(points)

	if err := p.Save(20.32*vg.Centimeter, 20.32*vg.Centimeter, "plot.png"); err != nil {
		log.Fatal(err)
	}

}
