package components

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BodyInput struct {
	Input     textarea.Model
	isFocused bool
}

func NewBodyInput() BodyInput {

	ta := textarea.New()

	ta.Placeholder = "Enter Body"
	ta.SetWidth(52)
	ta.SetHeight(5)
	ta.MaxHeight = 256

	return BodyInput{
		Input:     ta,
		isFocused: false,
	}
}

func (b *BodyInput) Focus() {
	b.isFocused = true
	b.Input.Focus()
}

func (b *BodyInput) Blur() {
	b.isFocused = false
	b.Input.Blur()
}

func (b BodyInput) IsFocused() bool {
	return b.isFocused
}

func (b *BodyInput) Update(msg tea.Msg) (BodyInput, tea.Cmd) {
	var cmd tea.Cmd

	if !b.IsFocused() {
		return *b, cmd
	}

	b.Input, cmd = b.Input.Update(msg)

	return *b, cmd

}

func (b BodyInput) View() string {
	bodyInputFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	bodyInputStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	if b.IsFocused() {
		return bodyInputFocusedStyle.Render(b.Input.View())
	}

	return bodyInputStyle.Render(b.Input.View())
}
