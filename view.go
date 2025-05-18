package main

import (
	"fmt"
	"strings"
)




func (m Model) View() string {
	
	status:= "Paused" 
	if m.Running {
			status = "Running"
	}
	minutes:= int(m.TimeLeft.Minutes())
	seconds:= int(m.TimeLeft.Seconds())%60
	timeLeftFormatted := fmt.Sprintf("%02d:%02d", minutes, seconds)

	
	    return fmt.Sprintf(
        "\nMode: %s\nTime Left: %s\nPomodoros Completed: %d\nStatus: %s\n\n%s\n\nPress 'q' to quit, 's' to start, 't' to toggle timer, 'r' to reset.",
        m.Mode, timeLeftFormatted, m.PomodoroCount, status, m.Progress.View(),
    )
	// return fmt.Sprintf(
	// 	"%s | %s | %02d:%02d\n\n[s]tart [t]oggle [r]eset [q]uit",
	// 	status, m.Mode.String(), minutes, seconds,
	// )
    
}
var blocks = []rune{' ', '▏', '▎', '▍', '▌', '▋', '▊', '▉', '█'}

func drawProgressBar(progress float64, width int) string {
    totalBlocks := progress * float64(width)
    fullBlocks := int(totalBlocks)
    partialBlock := int((totalBlocks - float64(fullBlocks)) * float64(len(blocks)-1))

    bar := strings.Repeat("█", fullBlocks)
    if fullBlocks < width {
        bar += string(blocks[partialBlock])
        bar += strings.Repeat(" ", width-fullBlocks-1)
    }
    return "[" + bar + "]"
}
