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
			Height: 300,
			Panels: []LayoutPanel{
				{
					Drawer: &plot.Lines{
						Labels: plot.Labels{
							Title: "Worker age classes",
							X:     "Age [d]",
							Y:     "Count",
						},
						YLim: [2]float64{0, 2000},
					},
					Observer: &obs.AgeStructure{},
				},
			},
		}, {
			Height: 260,
			Panels: []LayoutPanel{
				{
					Drawer: &plot.TimeSeries{
						Labels: plot.Labels{
							Title: "Worker cohorts",
							X:     "Time [d]",
							Y:     "Count",
						},
						MaxRows: 730,
					},
					Observer: &obs.WorkerCohorts{Cumulative: true},
				}, {
					Drawer: &plot.TimeSeries{
						Labels: plot.Labels{
							Title: "In-hive stores",
							X:     "Time [d]",
							Y:     "Amount [kg]",
						},
						Columns: []string{"Honey", "Pollen x20"},
						MaxRows: 730,
					},
					Observer: &obs.Stores{PollenFactor: 20},
				},
			},
		}, {
			Height: 300,
			Panels: []LayoutPanel{
				{
					Drawer: &plot.Lines{
						Labels: plot.Labels{
							Title: "Foraging activity",
							X:     "Time [rounds]",
							Y:     "Foragers",
						},
						X:    "Round",
						XLim: [2]float64{0, 42},
					},
					Observer: &obs.ForagingStats{},
				},
			},
		},
	},
}
