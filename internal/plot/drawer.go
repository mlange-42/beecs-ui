package plot

import (
	"github.com/mlange-42/ark/ecs"
	"gonum.org/v1/plot/vg/vgimg"
)

type Drawer interface {
	Initialize(w *ecs.World, observer any) error
	Update(w *ecs.World)
	SetChanged()
	Draw(world *ecs.World, canvas *vgimg.Canvas) bool
}
