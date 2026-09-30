package components

import "github.com/charmbracelet/lipgloss"

type Footer struct {
	Text string
}

func NewFooter() Footer {
	return Footer{
		Text: "[Tab] Switch Focus | [Enter] Send | [Right/Left] Change Method | [Ctrl+C] Quit",
	}
}

func (f Footer) View() string {

	footerStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		Padding(1, 2).
		Foreground(lipgloss.Color("#fff")).
		Bold(true).Width(54)

	return footerStyle.Render(f.Text)
}
