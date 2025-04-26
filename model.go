package main

import "time"
type Mode int

const (
	// Node types
	Pomodoro Mode = iota
	ShortBreak
	LongBreak
)
func (m Mode) String() string {
    switch m {
    case Pomodoro:
        return "Pomodoro"
    case ShortBreak:
        return "Short Break"
    case LongBreak:
        return "Long Break"
    default:
        return "Unknown"
    }
}
func (m Mode) SessionDuration() time.Duration {
    switch m {
    case Pomodoro:
        return 25 * time.Minute
    case ShortBreak:
        return 5 * time.Minute
    case LongBreak:
        return 15 * time.Minute
    default:
        return 0
    }
}
// Model holds the application state
type Model struct {
    Mode          Mode          // Current session type
    TimeLeft      time.Duration // Time remaining
    PomodoroCount int           // Number of completed Pomodoros
    Running       bool          // Is the timer active?
    LastTick     time.Time     // Last tick time
}


func initialModel() Model {
return Model{
	Mode:          Pomodoro,
	TimeLeft:      25 * time.Minute,
	PomodoroCount: 0,
	Running:       false,
	}
}


