package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)
type TickMsg struct{}

func (m Model) Init() tea.Cmd {

	return nil
}

func  tickFunc(t time.Time) tea.Msg {
	return TickMsg{}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			// Quit the app
			return m, tea.Quit
		case "s":
			// Start the timer
			if !m.Running {
				m.Running = true
			}
		case "t":
			// Toggle the timer (start or stop)
			m.Running = !m.Running
		case "r":
			// Reset the timer
			m.Running = false
			m.TimeLeft = 25 * time.Minute
		}

	case TickMsg:
		if m.Running {
			// Update the time left if the timer is running
			if time.Since(m.LastTick) >=time.Second  {
				if( m.TimeLeft > 0) {
				m.TimeLeft -= time.Second
				}
			} else {
				// Timer is done, transition to the next mode
				m.Running = false
				switch m.Mode {
				case Pomodoro:
					m.Mode = ShortBreak
					m.TimeLeft = 5 * time.Minute
				case ShortBreak:
					m.Mode = Pomodoro
					m.TimeLeft = 25 * time.Minute
					m.PomodoroCount++
				case LongBreak:
					m.Mode = Pomodoro
					m.TimeLeft = 25 * time.Minute
				}
			}
			m.LastTick = time.Now()
		}
	}
	return m, tea.Tick(time.Second,tickFunc )
}