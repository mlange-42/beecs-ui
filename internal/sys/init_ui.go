package sys

import (
	"io/fs"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/ui"
)

// InitUI system.
type InitUI struct {
	ResetFn    func(parameters map[string]any, speed uint8)
	LayoutData fs.FS
	Layout     string
	ui         ui.UI
}

// Initialize the system
func (s *InitUI) Initialize(world *ecs.World) {
	s.ui = ui.New(world, s.LayoutData, s.Layout, s.ResetFn)

	ecs.AddResource(world, &s.ui)
}

// Update the system
func (s *InitUI) Update(world *ecs.World) {}

// Finalize the system
func (s *InitUI) Finalize(world *ecs.World) {}
