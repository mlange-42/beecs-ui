package plot

import (
	"github.com/mlange-42/arche/ecs"
	"gonum.org/v1/plot/vg/vgimg"
)

type Drawer interface {
	Initialize(w *ecs.World)
	Update(w *ecs.World)
	Draw(world *ecs.World, canvas *vgimg.Canvas)
}
