package sys

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/beecs-ui/internal/res"
)

// GameControls system.
type GameControls struct {
	PauseKey      ebiten.Key
	StepKey       ebiten.Key
	SlowerKey     rune
	FasterKey     rune
	FullscreenKey ebiten.Key

	time      ecs.Resource[res.GameTick]
	speed     ecs.Resource[res.GameSpeed]
	prevSpeed uint8

	inputChars []rune
}

// Initialize the system
func (s *GameControls) Initialize(world *ecs.World) {
	s.time = ecs.NewResource[res.GameTick](world)
	s.speed = ecs.NewResource[res.GameSpeed](world)

	speed := s.speed.Get()

	ebiten.SetTPS(int(speed.Speeds[speed.SpeedIndex]))
}

// Update the system
func (s *GameControls) Update(world *ecs.World) {
	time := s.time.Get()
	speed := s.speed.Get()

	if inpututil.IsKeyJustPressed(s.FullscreenKey) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(s.PauseKey) {
		speed.Pause = !speed.Pause
		speed.NextPause = -1
	}
	if inpututil.IsKeyJustPressed(s.StepKey) {
		steps := 1
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			steps = 30
			if ebiten.IsKeyPressed(ebiten.KeyAlt) {
				steps = 365

			}
		}
		speed.Pause = false
		speed.NextPause = time.Tick + int64(steps)
	}

	if speed.NextPause == time.Tick {
		speed.Pause = true
		speed.NextPause = -1

	}

	s.inputChars = ebiten.AppendInputChars(s.inputChars)

	if speed.SpeedIndex > 0 && slices.Contains(s.inputChars, s.SlowerKey) {
		speed.SpeedIndex--
	}
	if int(speed.SpeedIndex) < len(speed.Speeds)-1 && slices.Contains(s.inputChars, s.FasterKey) {
		speed.SpeedIndex++
	}

	s.inputChars = s.inputChars[:0]

	if s.prevSpeed != speed.SpeedIndex {
		ebiten.SetTPS(int(speed.Speeds[speed.SpeedIndex]))
		s.prevSpeed = speed.SpeedIndex
	}
}

// Finalize the system
func (s *GameControls) Finalize(world *ecs.World) {}
