package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLInput struct {
	Input     textinput.Model
	isFocused bool
}

func NewURLInput() URLInput {
	ti := textinput.New()

	ti.Placeholder = "Enter URL"
	ti.CharLimit = 256
	ti.Width = 43

	return URLInput{
		Input:     ti,
		isFocused: false,
	}
}

// Focus() and Blur() both are mutate methods that mutate their own states so use pointer *URLInput
func (u *URLInput) Focus() {
	u.isFocused = true
	u.Input.Focus()
	// jugaad
	u.Input.Placeholder = ""
	u.Input.Width = 40
}

func (u *URLInput) Blur() {
	u.isFocused = false
	u.Input.Blur()
	// jugaad
	u.Input.Placeholder = "Enter URL"
	u.Input.Width = 43
}

// it is a read method so URLInput will also work as a receiver instead of *URLInput
func (u URLInput) IsFocused() bool {
	return u.isFocused
}

func (u *URLInput) Update(msg tea.Msg) (URLInput, tea.Cmd) {

	var cmd tea.Cmd

	if !u.IsFocused() {
		return *u, cmd
	}

	u.Input, cmd = u.Input.Update(msg)

	return *u, cmd
}

func (u URLInput) View() string {

	urlInputFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	if u.IsFocused() {
		return urlInputFocusedStyle.Render(u.Input.View())
	}

	return urlInputStyle.Render(u.Input.View())
}
