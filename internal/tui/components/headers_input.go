package components

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HeadersInput struct {
	Headers        []HeadersRow
	isFocused      bool
	ActiveRowIndex int
}

func (h *HeadersInput) Focus() {
	h.isFocused = true
	h.Headers[h.ActiveRowIndex].FocusRow() // for initial focus
}

func (h *HeadersInput) Blur() {
	h.isFocused = false
	// when global focus will shift from headers input to send btn,
	// this Blur() will be called in FocusNext() func as currComp.Blur()
	// so with this func call, this logic will explicitily blur the last focused row
	if len(h.Headers) > 0 && h.ActiveRowIndex >= 0 && h.ActiveRowIndex < len(h.Headers) {
		h.Headers[h.ActiveRowIndex].BlurRow()
	}
}

func (h HeadersInput) IsFocused() bool {
	return h.isFocused
}

func (h *HeadersInput) FocusPrevRow() {
	totalRows := len(h.Headers)
	if totalRows == 0 || totalRows == 1 {
		return
	}

	if h.ActiveRowIndex >= 0 && h.ActiveRowIndex < totalRows {
		h.Headers[h.ActiveRowIndex].BlurRow()
	}

	h.ActiveRowIndex = (h.ActiveRowIndex - 1 + totalRows) % totalRows
	h.Headers[h.ActiveRowIndex].FocusRow()
}

func (h *HeadersInput) FocusNextRow() {
	totalRows := len(h.Headers)
	if totalRows == 0 || totalRows == 1 {
		return
	}

	if h.ActiveRowIndex >= 0 && h.ActiveRowIndex < totalRows {
		h.Headers[h.ActiveRowIndex].BlurRow()
	}

	h.ActiveRowIndex = (h.ActiveRowIndex + 1) % totalRows
	h.Headers[h.ActiveRowIndex].FocusRow()
}

func NewHeadersInput() HeadersInput {

	var headers []HeadersRow

	for i := 0; i < 3; i++ {
		headers = append(headers, NewHeadersRow())
	}

	return HeadersInput{
		Headers:        headers,
		isFocused:      false,
		ActiveRowIndex: 0,
	}
}

func (h HeadersInput) GetCurrRowIndex() int {
	return h.ActiveRowIndex
}

func (h HeadersInput) GetCurrRow() HeadersRow {
	return h.Headers[h.ActiveRowIndex]
}

func (h *HeadersInput) AddRow() {
	newRow := NewHeadersRow()
	h.Headers = append(h.Headers, newRow)
}

func (h *HeadersInput) RemoveRow() {
	// there should be at least one row
	if len(h.Headers) == 1 {
		return
	}

	rowToRemove := h.ActiveRowIndex
	headers := slices.Delete(h.Headers, rowToRemove, rowToRemove+1)
	h.Headers = headers
	// after removing row, explicitly decrease active row index to avoid panic of index out of range
	if h.ActiveRowIndex >= len(h.Headers) && len(h.Headers) > 0 {
		h.ActiveRowIndex = len(h.Headers) - 1
	}
	// after removing row, explicitly focus on the last/new row
	h.Headers[h.ActiveRowIndex].FocusRow()
}

func (h *HeadersInput) Update(msg tea.Msg) (HeadersInput, tea.Cmd) {
	var cmd tea.Cmd

	if !h.IsFocused() {
		h.Headers[h.ActiveRowIndex].BlurRow() // blur current focused row if component is not focused
		return *h, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+a":
			h.AddRow()
		case "ctrl+r":
			h.RemoveRow()
		case "down":
			h.FocusNextRow()
			return *h, cmd
		case "up":
			h.FocusPrevRow()
			return *h, cmd
		}
	}

	h.Headers[h.ActiveRowIndex], cmd = h.Headers[h.ActiveRowIndex].Update(msg)

	return *h, cmd

}

func (h HeadersInput) View() string {

	headersInputFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	headersInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	tagStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1)

	tag := "Headers: "

	headers := []string{tagStyle.Render(tag)}

	for i := range h.Headers {
		headers = append(headers, h.Headers[i].View())
	}

	var headersLayout = lipgloss.JoinVertical(
		lipgloss.Left,
		headers...,
	)

	if h.isFocused {
		return headersInputFocusedStyle.Render(headersLayout)
	}

	return headersInputStyle.Render(headersLayout)

}
