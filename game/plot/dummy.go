package plot

import (
	"github.com/mlange-42/arche/ecs"
	"gonum.org/v1/plot/vg/vgimg"
)

type Dummy struct{}

// Initialize the drawer.
func (t *Dummy) Initialize(w *ecs.World) {}

// Update the drawer.
func (t *Dummy) Update(w *ecs.World) {}

// Draw the drawer.
func (t *Dummy) Draw(w *ecs.World, canvas *vgimg.Canvas) {}
