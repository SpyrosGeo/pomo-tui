package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


func main(){
	    p := tea.NewProgram(initialModel())
    if err := p.Start(); err != nil {
        fmt.Println("Error starting app:", err)
        os.Exit(1)
    }
}