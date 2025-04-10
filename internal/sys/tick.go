package sys

import (
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/res"
)

// Tick system.
type Tick struct {
	speed ecs.Resource[res.GameSpeed]
	time  ecs.Resource[res.GameTick]
}

// Initialize the system
func (s *Tick) Initialize(world *ecs.World) {
	s.speed = ecs.NewResource[res.GameSpeed](world)
	s.time = ecs.NewResource[res.GameTick](world)
}

// Update the system
func (s *Tick) Update(world *ecs.World) {
	if !s.speed.Get().Pause {
		s.time.Get().Tick++
	}
}

// Finalize the system
func (s *Tick) Finalize(world *ecs.World) {}
