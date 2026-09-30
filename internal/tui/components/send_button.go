package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SendButton struct {
	onClick   func()
	isFocused bool
}

func NewSendButton(onclick func()) SendButton {
	return SendButton{
		onClick:   onclick,
		isFocused: false,
	}
}

func (s *SendButton) Focus() {
	s.isFocused = true
}

func (s *SendButton) Blur() {
	s.isFocused = false
}

func (s SendButton) IsFocused() bool {
	return s.isFocused
}

func (s *SendButton) Update(msg tea.Msg) (SendButton, tea.Cmd) {
	var cmd tea.Cmd

	if !s.IsFocused() {
		return *s, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			s.onClick()
		}
	}

	return *s, cmd
}

func (s SendButton) View() string {

	buttonStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1).Margin(0, 1)

	buttonFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1).
		Bold(true).Margin(0, 1).Foreground(lipgloss.Color("#fff"))

	if s.IsFocused() {
		return buttonFocusedStyle.Render("Send")
	}

	return buttonStyle.Render("Send")

}
