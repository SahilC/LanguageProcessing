package main

import (
    "math"
    "github.com/gonum/plot"
    "github.com/gonum/plot/plotter"
    "github.com/gonum/plot/plotutil"
    "github.com/gonum/plot/vg"
)
func ProcessPoint(vals []float64) plotter.XYs {
    pts := make(plotter.XYs, len(vals))
    for i, k := range vals {
        pts[i].X = math.Log(float64(i+1))
        pts[i].Y = math.Log(k+1)
    }
    return pts
}

func PlotLogLog(vals [][]float64,filename string, labels []string) {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }

    p.Title.Text = "Log-Log plot of Frequency v/s Rank"
    p.X.Label.Text = "Rank"
    p.Y.Label.Text = "Frequency"
    // for t := 0; t< len(vals);t++ {
    //     pts := make(plotter.XYs, len(vals[t]))
    //     for i, k := range vals[t] {
    //         pts[i].X = math.Log(float64(i+1))
    //         pts[i].Y = math.Log(k+1)
    //     }
    //
    // }
    err = plotutil.AddLinePoints(p,"Sherlock Holmes", ProcessPoint(vals[0]),
                                    "Alice in Wonderland",ProcessPoint(vals[1]),
                                    "Shakespear's works", ProcessPoint(vals[2]),
                                    "Persuasion", ProcessPoint(vals[3]),
                                    "Siddaratha", ProcessPoint(vals[4]),
                                    "Twitter", ProcessPoint(vals[5]))
    if err != nil {
        panic(err)
    }
    // Save the plot to a PNG file.
    if err := p.Save(5*vg.Inch, 6*vg.Inch, filename+".png"); err != nil {
        panic(err)
    }
}
