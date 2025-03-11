package plot

import (
	"fmt"
	"math"

	"github.com/mlange-42/ark-tools/observer"
	"github.com/mlange-42/ark/ecs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// Lines plot drawer.
//
// Creates a line series per column of the observer.
// Replaces the complete data by the table provided by the observer on every update.
// Particularly useful for live histograms.
type Lines struct {
	observer     observer.Table // Observer providing a data series for lines.
	X            string         // X column name. Optional. Defaults to row index.
	Y            []string       // Y column names. Optional. Defaults to all but X column.
	XLim         [2]float64     // X axis limits. Optional, default auto.
	YLim         [2]float64     // Y axis limits. Optional, default auto.
	Labels       Labels         // Labels for plot and axes. Optional.
	DrawInterval int

	xIndex   int
	yIndices []int

	headers []string
	series  []plotter.XYs

	drawStep   uint64
	hasChanged bool
}

// Initialize the drawer.
func (l *Lines) Initialize(w *ecs.World, obs any) error {
	table, castOk := obs.(observer.Table)
	if !castOk {
		return fmt.Errorf("unable to cast %T to table observer", obs)
	}
	l.observer = table
	l.observer.Initialize(w)

	l.headers = l.observer.Header()

	var ok bool
	if l.X == "" {
		l.xIndex = -1
	} else {
		l.xIndex, ok = find(l.headers, l.X)
		if !ok {
			return fmt.Errorf("x column '%s' not found", l.X)
		}
	}

	if len(l.Y) == 0 {
		l.yIndices = make([]int, 0, len(l.headers))
		for i := 0; i < len(l.headers); i++ {
			if i != l.xIndex {
				l.yIndices = append(l.yIndices, i)
			}
		}
	} else {
		l.yIndices = make([]int, len(l.Y))
		for i, y := range l.Y {
			l.yIndices[i], ok = find(l.headers, y)
			if !ok {
				return fmt.Errorf("y column '%s' not found", y)
			}
		}
	}

	l.series = make([]plotter.XYs, len(l.yIndices))

	return nil
}

func (l *Lines) SetChanged() {
	l.hasChanged = true
}

// Update the drawer.
func (l *Lines) Update(w *ecs.World) {
	l.observer.Update(w)
	l.hasChanged = true
}

// Draw the drawer.
func (l *Lines) Draw(w *ecs.World, canvas *vgimg.Canvas) bool {
	if !l.hasChanged || (l.DrawInterval > 1 && l.drawStep%uint64(l.DrawInterval) != 0) {
		l.drawStep++
		return false
	}

	l.updateData(w)

	p := plot.New()
	setLabels(p, l.Labels)

	p.X.Tick.Marker = removeLastTicks{}

	if l.YLim[0] != 0 || l.YLim[1] != 0 {
		p.Y.Min = l.YLim[0]
		p.Y.Max = l.YLim[1]
	}

	if l.XLim[0] != 0 || l.XLim[1] != 0 {
		p.X.Min = l.XLim[0]
		p.X.Max = l.XLim[1]
	}

	p.Legend = plot.NewLegend()
	p.Legend.TextStyle.Font.Variant = "Mono"

	for i := 0; i < len(l.series); i++ {
		idx := l.yIndices[i]
		lines, err := plotter.NewLine(l.series[i])
		if err != nil {
			panic(err)
		}
		lines.Color = defaultColors[i%len(defaultColors)]
		p.Add(lines)
		p.Legend.Add(l.headers[idx], lines)
	}

	p.Draw(draw.New(canvas))

	l.drawStep++
	l.hasChanged = false
	return true
}

func (l *Lines) updateData(w *ecs.World) {
	data := l.observer.Values(w)
	xi := l.xIndex
	yis := l.yIndices

	for i, idx := range yis {
		l.series[i] = l.series[i][:0]
		for j, row := range data {
			x := float64(j)
			if xi >= 0 {
				x = row[xi]
			}
			if math.IsNaN(row[idx]) {
				continue
			}
			l.series[i] = append(l.series[i], plotter.XY{X: x, Y: row[idx]})
		}
	}
}
