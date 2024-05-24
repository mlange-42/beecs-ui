package ui

import (
	"github.com/mlange-42/beecs-ui/game/plot"
	"github.com/mlange-42/beecs/obs"
)

var defaultLayout = Layout{
	Parameters: []ParameterSection{
		{
			Text: "Initialization",
			Parameters: []Parameter{
				{
					Parameter: "params.InitialPopulation.Count",
					SliderInt: &SliderInt{
						Min: 1_000,
						Max: 50_000,
					},
				},
			},
		},
		{
			Text: "HandlingTime",
			Parameters: []Parameter{
				{
					Parameter: "params.HandlingTime.ConstantHandlingTime",
					Toggle:    &Toggle{},
				},
				{
					Parameter: "params.HandlingTime.NectarGathering",
					SliderFloat: &SliderFloat{
						Min:       60,
						Max:       3600,
						Precision: 1,
					},
				},
				{
					Parameter: "params.HandlingTime.PollenGathering",
					SliderFloat: &SliderFloat{
						Min:       60,
						Max:       3600,
						Precision: 1,
					},
				},
			},
		},
	},
	Rows: []LayoutRow{
		{
			Height: 350,
			Panels: []LayoutPanel{
				{
					Drawer: &plot.TimeSeries{
						Observer: &obs.WorkerCohorts{},
						Labels: plot.Labels{
							Title: "Worker cohorts",
							X:     "Time [d]",
							Y:     "Count",
						},
					},
				},
				{Drawer: &plot.Dummy{}},
			},
		}, {
			Height: 200,
			Panels: []LayoutPanel{
				{Drawer: &plot.Dummy{}}, {Drawer: &plot.Dummy{}}, {Drawer: &plot.Dummy{}},
			},
		}, {
			Height: 200,
			Panels: []LayoutPanel{
				{Drawer: &plot.Dummy{}}, {Drawer: &plot.Dummy{}}, {Drawer: &plot.Dummy{}},
			},
		}, {
			Height: 400,
			Panels: []LayoutPanel{
				{Drawer: &plot.Dummy{}},
			},
		},
	},
}
