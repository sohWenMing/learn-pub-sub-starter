package routing

import "time"

type PlayingState struct {
	IsPaused bool
}

// the gameLog just keeps track of the time of log, the message and the username
type GameLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}
