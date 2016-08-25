package main

import (
    "math"
    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/plotutil"
    "github.com/gonum/plot/vg"
)

func PlotLogLog(vals []float64) {
    pts := make(plotter.XYs, len(vals))
    for i, k := range vals {
        pts[i].X = math.Log(float64(i+1))
        pts[i].Y = math.Log(k+1)
    }

    p, err := plot.New()
    if err != nil {
        panic(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    err = plotutil.AddLinePoints(p,
        "First", pts)
    if err != nil {
        panic(err)
    }

    // Save the plot to a PNG file.
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
        panic(err)
    }
}
