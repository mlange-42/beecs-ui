package sys

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/beecs-ui/game/res"
)

// RenderUI is a system to render the user interface.
type RenderUI struct {
	screen generic.Resource[res.Screen]
	ui     generic.Resource[res.UI]
}

// InitializeUI the system
func (s *RenderUI) InitializeUI(world *ecs.World) {
	s.ui = generic.NewResource[res.UI](world)
	s.screen = generic.NewResource[res.Screen](world)
}

// UpdateUI the system
func (s *RenderUI) UpdateUI(world *ecs.World) {
	screen := s.screen.Get()
	ui := s.ui.Get()

	ui.Draw(screen.Image)
}

// PostUpdateUI the system
func (s *RenderUI) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *RenderUI) FinalizeUI(world *ecs.World) {}
