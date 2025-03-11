package sys

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/res"
	"github.com/mlange-42/beecs-ui/internal/ui"
)

// RenderUI is a system to render the user interface.
type RenderUI struct {
	screen ecs.Resource[res.Screen]
	ui     ecs.Resource[ui.UI]
}

// InitializeUI the system
func (s *RenderUI) InitializeUI(world *ecs.World) {
	s.ui = ecs.NewResource[ui.UI](world)
	s.screen = ecs.NewResource[res.Screen](world)
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
