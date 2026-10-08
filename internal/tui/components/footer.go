package components

import "github.com/charmbracelet/lipgloss"

type Footer struct {
	Text string
}

func NewFooter() Footer {
	return Footer{
		Text: `[Tab]        Switch global focus
[Enter]      Send or switch key/value focus
[Right/Left] Change HTTP method
[Shift+Tab]  Switch tabs
[Up/Down]    Switch header/param rows focus
[Ctrl+A]     Add new header/param row
[Ctrl+R]     Remove current header/param row
[Ctrl+C]     Quit`,
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
