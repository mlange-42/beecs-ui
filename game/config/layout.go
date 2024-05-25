package config

import "github.com/mlange-42/beecs-ui/game/plot"

type Layout struct {
	Parameters []ParameterSection
	Rows       []LayoutRow
}

type LayoutRow struct {
	Height int
	Panels []LayoutPanel
}

type LayoutPanel struct {
	Drawer   plot.Drawer
	Observer any
}
