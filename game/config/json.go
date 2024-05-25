package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"reflect"

	"github.com/mlange-42/beecs-ui/game/plot"
	"github.com/mlange-42/beecs-ui/registry"
)

func FromFile(f fs.FS, path string, templatesFS fs.FS, paramTemplates string, panelTemplates string) (*Layout, error) {
	pars, panels, err := readTemplates(templatesFS, paramTemplates, panelTemplates)
	if err != nil {
		return nil, err
	}

	content, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	layoutJs := layoutJs{}
	if err := decoder.Decode(&layoutJs); err != nil {
		return nil, err
	}

	layout := &Layout{
		Parameters: layoutJs.Parameters,
	}

	for i := range layout.Parameters {
		section := &layout.Parameters[i]
		for j := range section.Parameters {
			tpl := section.Parameters[j].Template
			if tpl != "" {
				replace, ok := pars[tpl]
				if !ok {
					return nil, fmt.Errorf("no parameter template '%s' found", tpl)
				}
				section.Parameters[j] = replace
			}
		}
	}

	for i, row := range layoutJs.Rows {
		layout.Rows = append(layout.Rows, LayoutRow{
			Height: row.Height,
		})
		for _, panel := range row.Panels {
			if panel.Template != "" {
				replace, ok := panels[panel.Template]
				if !ok {
					return nil, fmt.Errorf("no panel template '%s' found", panel.Template)
				}
				panel = replace
			}

			drawer, err := decodeDrawer(panel.Drawer, panel.DrawerConfig)
			if err != nil {
				return nil, err
			}
			observer, err := decodeObserver(panel.Observer, panel.ObserverConfig)
			if err != nil {
				return nil, err
			}
			layout.Rows[i].Panels = append(layout.Rows[i].Panels, LayoutPanel{
				Drawer:   drawer,
				Observer: observer,
			})
		}
	}

	return layout, nil
}

func readTemplates(f fs.FS, paramTemplates string, panelTemplates string) (map[string]Parameter, map[string]layoutPanelJs, error) {
	pars := map[string]Parameter{}
	panels := map[string]layoutPanelJs{}

	if f == nil {
		return pars, panels, nil
	}

	content, err := fs.ReadFile(f, paramTemplates)
	if err != nil {
		return nil, nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&pars); err != nil {
		return nil, nil, err
	}

	content, err = fs.ReadFile(f, panelTemplates)
	if err != nil {
		return nil, nil, err
	}
	decoder = json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&panels); err != nil {
		return nil, nil, err
	}

	return pars, panels, nil
}

func decodeObserver(name string, config entry) (any, error) {
	tp, ok := registry.GetObserver(name)
	if !ok {
		return nil, fmt.Errorf("observer type '%s' is not registered", name)
	}
	observerVal := reflect.New(tp).Interface()
	if len(config.Bytes) == 0 {
		config.Bytes = []byte("{}")
	}

	decoder := json.NewDecoder(bytes.NewReader(config.Bytes))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&observerVal); err != nil {
		return nil, err
	}
	return observerVal, nil
}

func decodeDrawer(name string, config entry) (plot.Drawer, error) {
	tp, ok := registry.GetDrawer(name)
	if !ok {
		return nil, fmt.Errorf("drawer type '%s' is not registered", name)
	}
	drawerVal := reflect.New(tp).Interface()
	if len(config.Bytes) == 0 {
		config.Bytes = []byte("{}")
	}

	decoder := json.NewDecoder(bytes.NewReader(config.Bytes))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&drawerVal); err != nil {
		return nil, err
	}

	drawerCast, ok := drawerVal.(plot.Drawer)
	if !ok {
		return nil, fmt.Errorf("type '%s' is not a plot.Drawer", tp.String())
	}

	return drawerCast, nil
}
