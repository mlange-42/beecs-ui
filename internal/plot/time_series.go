package plot

import (
	"fmt"

	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// TimeSeries plot drawer.
//
// Creates a line series per column of the observer.
// Adds one row to the data per update.
type TimeSeries struct {
	observer       observer.Row // Observer providing a data row per update.
	Columns        []string     // Columns to show, by name. Optional, default all.
	UpdateInterval int          // Interval for getting data from the the observer, in model ticks. Optional.
	DrawInterval   int
	Labels         Labels // Labels for plot and axes. Optional.
	MaxRows        int    // Maximum number of rows to keep. Zero means unlimited. Optional.

	indices    []int
	headers    []string
	series     []plotter.XYs
	scale      float64
	step       int64
	drawStep   uint64
	hasChanged bool
}

// append a y value to each series, with a common x value.
func (t *TimeSeries) append(x float64, values []float64) {
	for i := 0; i < len(t.series); i++ {
		t.series[i] = append(t.series[i], plotter.XY{X: x, Y: values[i]})
		if t.MaxRows > 0 && len(t.series[i]) > t.MaxRows {
			t.series[i] = t.series[i][len(t.series[i])-t.MaxRows:]
		}
	}
}

// Initialize the drawer.
func (t *TimeSeries) Initialize(w *ecs.World, obs any) error {
	row, ok := obs.(observer.Row)
	if !ok {
		return fmt.Errorf("unable to cast %T to row observer", obs)
	}
	t.observer = row

	t.observer.Initialize(w)

	t.headers = t.observer.Header()

	if len(t.Columns) == 0 {
		t.indices = make([]int, len(t.headers))
		for i := 0; i < len(t.indices); i++ {
			t.indices[i] = i
		}
	} else {
		t.indices = make([]int, len(t.Columns))
		var ok bool
		for i := 0; i < len(t.indices); i++ {
			t.indices[i], ok = find(t.headers, t.Columns[i])
			if !ok {
				return fmt.Errorf("column '%s' not found", t.Columns[i])
			}
		}
	}

	t.series = make([]plotter.XYs, len(t.headers))

	t.scale = calcScaleCorrection()
	t.step = 0

	return nil
}

// Update the drawer.
func (t *TimeSeries) Update(w *ecs.World) {
	t.observer.Update(w)
	if t.UpdateInterval <= 1 || t.step%int64(t.UpdateInterval) == 0 {
		t.append(float64(t.step), t.observer.Values(w))
		t.hasChanged = true
	}
	t.step++
}

func (t *TimeSeries) SetChanged() {
	t.hasChanged = true
}

// Draw the drawer.
func (t *TimeSeries) Draw(w *ecs.World, canvas *vgimg.Canvas) bool {
	if !t.hasChanged || (t.DrawInterval > 1 && t.drawStep%uint64(t.DrawInterval) != 0) {
		t.drawStep++
		return false
	}

	p := plot.New()
	setLabels(p, t.Labels)

	p.X.Tick.Marker = removeLastTicks{}

	p.Legend = plot.NewLegend()
	p.Legend.TextStyle.Font.Variant = "Mono"

	for i, idx := range t.indices {
		lines, err := plotter.NewLine(t.series[idx])
		if err != nil {
			panic(err)
		}
		lines.Color = defaultColors[i%len(defaultColors)]
		p.Add(lines)
		p.Legend.Add(t.headers[idx], lines)
	}

	p.Draw(draw.New(canvas))

	t.drawStep++
	t.hasChanged = false
	return true
}
