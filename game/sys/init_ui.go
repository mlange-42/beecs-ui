package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
)

// InitUI system.
type InitUI struct {
	ui res.UI
}

// Initialize the system
func (s *InitUI) Initialize(world *ecs.World) {
	s.ui = res.NewUI(world)

	ecs.AddResource(world, &s.ui)
}

// Update the system
func (s *InitUI) Update(world *ecs.World) {}

// Finalize the system
func (s *InitUI) Finalize(world *ecs.World) {}
