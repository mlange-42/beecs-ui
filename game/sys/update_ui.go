package sys

import (
	"fmt"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
	"github.com/mlange-42/beecs-ui/game/ui"
)

// UpdateUI system.
type UpdateUI struct {
	ui   *ui.UI
	time *res.GameTick
}

// Initialize the system
func (s *UpdateUI) Initialize(world *ecs.World) {
	s.ui = ecs.GetResource[ui.UI](world)
	s.time = ecs.GetResource[res.GameTick](world)
}

// Update the system
func (s *UpdateUI) Update(world *ecs.World) {
	s.ui.InfoLabel.Label = fmt.Sprintf("Tick %5d", s.time.Tick)

	s.ui.Update()
}

// Finalize the system
func (s *UpdateUI) Finalize(world *ecs.World) {}
