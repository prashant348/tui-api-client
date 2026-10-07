package components

import "github.com/charmbracelet/lipgloss"

type Footer struct {
	Text string
}

func NewFooter() Footer {
	return Footer{
		Text: `[Tab]        Switch global focus
[Enter]      Send or Switch Key/Value focus
[Right/Left] Change method
[Ctrl+C]     Quit
[Up/Down]    Switch header rows focus
[Ctrl+A]     Add new header row
[Ctrl+R]     Remove current header row`,
	}
}

func (f Footer) View() string {

	footerStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		Padding(1, 2).
		Foreground(lipgloss.Color("#fff")).
		Bold(true).Width(54)

	return footerStyle.Render(f.Text)
}
