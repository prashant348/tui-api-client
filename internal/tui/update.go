package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// check msg type
	switch msg := msg.(type) {
	// if msg type is tea.KeyMsg
	case tea.KeyMsg:
		// since msg type is tea.KeyMsg, check pressed key value via msg.String()
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			if m.FocusIndex == 0 {
				m.FocusIndex = 1
				m.UrlInput.Focus()
			} else {
				m.FocusIndex = 0
				m.UrlInput.Blur()
			}

		case "left", "h":
			if m.FocusIndex == 0 && m.SelectedMethod > 0 {
				m.SelectedMethod--
			}
		case "right", "l":
			if m.FocusIndex == 0 && m.SelectedMethod < len(Methods)-1 {
				m.SelectedMethod++
			}

		case "ctrl+s", "enter":
			if m.UrlInput.Value() != "" {
				m.ResponseBody = fmt.Sprintf("Sending %s request to: %s...", Methods[m.SelectedMethod], m.UrlInput.Value())
			}

		}
	}
	// Agar URL bar focused hai, toh text input component ko update bhejo (typing handle karne ke liye)
	if m.FocusIndex == 1 {
		m.UrlInput, cmd = m.UrlInput.Update(msg)
	}

	return m, cmd
}
