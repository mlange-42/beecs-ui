package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
	"github.com/mlange-42/beecs-ui/game/ui"
)

// InitUI system.
type InitUI struct {
	ui ui.UI
}

// Initialize the system
func (s *InitUI) Initialize(world *ecs.World) {
	s.ui = ui.New(world, ecs.GetResource[res.Fonts](world), ecs.GetResource[res.Sprites](world))

	ecs.AddResource(world, &s.ui)
}

// Update the system
func (s *InitUI) Update(world *ecs.World) {}

// Finalize the system
func (s *InitUI) Finalize(world *ecs.World) {}
