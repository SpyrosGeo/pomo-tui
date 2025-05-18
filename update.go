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
    var cmds []tea.Cmd // Collect commands

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
                m.LastTick = time.Now() // Ensure LastTick is set when starting
            }
        case "t":
            // Toggle the timer (start or stop)
            m.Running = !m.Running
            if m.Running {
                m.LastTick = time.Now() // Set LastTick when resuming
            }
        case "r":
            // Reset the timer
            m.Running = false
            m.TimeLeft = 25 * time.Minute
            m.Progress.SetPercent(0) // Reset progress
        }

    case TickMsg:
        if m.Running {
            // Calculate elapsed time since the last tick
            elapsed := time.Since(m.LastTick)
            m.LastTick = time.Now()

            // Reduce the time left
            if m.TimeLeft > 0 {
                m.TimeLeft -= elapsed
                if m.TimeLeft < 0 {
                    m.TimeLeft = 0
                }
            }

            // Calculate progress percentage
            totalDuration := m.Mode.SessionDuration()
            progressPercent := 1 - m.TimeLeft.Seconds()/totalDuration.Seconds()
            if progressPercent < 0 {
                progressPercent = 0
            } else if progressPercent > 1 {
                progressPercent = 1
            }

            // Update progress bar
            progCmd := m.Progress.SetPercent(progressPercent)
            cmds = append(cmds, progCmd)

            // Handle timer completion
            if m.TimeLeft <= 0 {
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
                // Reset progress for new session
                cmds = append(cmds, m.Progress.SetPercent(0))
            }
        }
    }

    // Ensure the Tick command is always active if running
    if m.Running {
        cmds = append(cmds, tea.Tick(30*time.Millisecond, tickFunc))
    }

    // Return model with all commands
    return m, tea.Batch(cmds...)
}
