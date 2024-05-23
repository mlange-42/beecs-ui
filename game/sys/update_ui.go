package sys

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-ui/game/res"
	"github.com/mlange-42/beecs-ui/game/ui"
)

// UpdateUI system.
type UpdateUI struct {
	ui    *ui.UI
	time  *res.GameTick
	speed *res.GameSpeed
}

// Initialize the system
func (s *UpdateUI) Initialize(world *ecs.World) {
	s.ui = ecs.GetResource[ui.UI](world)
	s.time = ecs.GetResource[res.GameTick](world)
	s.speed = ecs.GetResource[res.GameSpeed](world)
}

// Update the system
func (s *UpdateUI) Update(world *ecs.World) {
	tick := s.time.Tick

	date := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * time.Duration(tick))

	s.ui.InfoLabel.Label = fmt.Sprintf("Tick %5d  %s", tick, date.Format("Jan _2 06"))

	if s.speed.Pause {
		s.ui.PauseButton.Text().Label = ">>"
	} else {
		s.ui.PauseButton.Text().Label = "||"
	}

	s.ui.Update()
}

// Finalize the system
func (s *UpdateUI) Finalize(world *ecs.World) {}
