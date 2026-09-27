package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	methodStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#10B98B")).
			Padding(0, 1)

	selectedMethodStyle = lipgloss.NewStyle().
				Bold(true).
				Background(lipgloss.Color("#6366F1")).
				Foreground(lipgloss.Color("#FFFFFF")).
				Padding(0, 1)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4B5563")).
			Padding(1, 2)
)

func (m Model) View() string {
	// 1. render method selector
	methodView := ""
	for i, method := range Methods {
		if i == m.SelectedMethod {
			methodView += selectedMethodStyle.Render(method) + " "
		} else {
			methodView += methodStyle.Render(method) + " "
		}
	}

	// 2. render header (method + url input)
	header := fmt.Sprintf(
		"Method: %s\nURL: %s",
		methodView,
		m.UrlInput.View(),
	)

	// 3. render response area
	responseView := boxStyle.Render(m.ResponseBody)

	// 4. Help text Footer
	footer := "\n[Tab] Switch Focus  |  [Left/Right] Change Method  |  [Enter] Send  |  [Ctrl+C] Quit"

	// 5. combine evenrything
	return header + "\n\n" + responseView + footer
}
