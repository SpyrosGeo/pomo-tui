package main

import (
	"fmt"
	"time"
)




func (m Model) View() string {
	   var status string
	   if m.Running {
		   status = "Running"
	   } else {
		   status = "Paused"
	   }
	   return fmt.Sprintf(
		"\nMode: %s\nTime Left: %v\nPomodoros Completed: %d\nStatus: %s\n\nPress 'q' to quit, 's' to start, 't' to toggle timer, 'r' to reset.",
		m.Mode, m.TimeLeft.Truncate(time.Second), m.PomodoroCount, status)
    
}