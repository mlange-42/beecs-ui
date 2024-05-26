package config

import "github.com/mlange-42/beecs-ui/internal/plot"

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

type layoutJs struct {
	Parameters []ParameterSection
	Rows       []layoutRowJs
}

type layoutRowJs struct {
	Height int
	Panels []layoutPanelJs
}

type layoutPanelJs struct {
	Template       string
	Drawer         string
	DrawerConfig   entry
	Observer       string
	ObserverConfig entry
}

type entry struct {
	Bytes []byte
}

func (e *entry) UnmarshalJSON(jsonData []byte) error {
	e.Bytes = jsonData
	return nil
}

func (e entry) MarshalJSON() ([]byte, error) {
	return e.Bytes, nil
}
