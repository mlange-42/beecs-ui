package res

// GameSpeed resource.
type GameSpeed struct {
	// Is the game paused?
	Pause bool
	// Next tick to pause
	NextPause int64
	// Speed index in Speeds.
	SpeedIndex uint8
	// Available speeds in TPS.
	Speeds []uint16
}

// GameTick resource.
type GameTick struct {
	// Current update tick. Stops when the game is paused.
	Tick int64
	// Current render tick. Does not stop when the game is paused.
	RenderTick int64
}
