package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HeadersRow struct {
	KeyInput       textinput.Model
	ValueInput     textinput.Model
	isFocused      bool
	isKeyFocused   bool
	isValueFocused bool
}

func (h *HeadersRow) FocusRow() {
	h.isFocused = true
	h.FocusRowKey()  // default focus is key input
	h.BlurRowValue() // and explicitily blur value input when key input is focused
}

func (h *HeadersRow) BlurRow() {
	h.isFocused = false
	h.BlurRowKey()
	h.BlurRowValue()
}

func (h HeadersRow) IsFocusedRow() bool {
	return h.isFocused
}

func (h *HeadersRow) FocusRowKey() {
	h.isKeyFocused = true
	h.KeyInput.Focus()
}

func (h *HeadersRow) FocusRowValue() {
	h.isValueFocused = true
	h.ValueInput.Focus()
}

func (h *HeadersRow) BlurRowKey() {
	h.isKeyFocused = false
	h.KeyInput.Blur()
}

func (h *HeadersRow) BlurRowValue() {
	h.isValueFocused = false
	h.ValueInput.Blur()
}

func NewHeadersRow() HeadersRow {

	ki := textinput.New()

	ki.Placeholder = "Key"
	ki.CharLimit = 256
	ki.Width = 24

	vi := textinput.New()

	vi.Placeholder = "Value"
	vi.CharLimit = 256
	vi.Width = 24

	headersRow := HeadersRow{
		KeyInput:   ki,
		ValueInput: vi,
		isFocused:  false,
	}

	return headersRow
}

func (h *HeadersRow) Update(msg tea.Msg) (HeadersRow, tea.Cmd) {
	var cmd tea.Cmd
	if !h.IsFocusedRow() {
		return *h, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if h.isKeyFocused {
				h.BlurRowKey()
				h.FocusRowValue()
			} else if h.isValueFocused {
				h.BlurRowValue()
				h.FocusRowKey()
			}
			// fmt.Printf("key focus: %v, value focus: %v", h.isKeyFocused, h.isValueFocused)
		}
	}

	if h.isKeyFocused {
		h.KeyInput, cmd = h.KeyInput.Update(msg)
	} else if h.isValueFocused {
		h.ValueInput, cmd = h.ValueInput.Update(msg)
	}

	return *h, cmd
}

func (h HeadersRow) View() string {

	headerRowStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	headerRowFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	headerRowLayout := lipgloss.JoinHorizontal(
		lipgloss.Center,
		h.KeyInput.View(),
		h.ValueInput.View(),
	)

	if h.IsFocusedRow() {
		return headerRowFocusedStyle.Render(headerRowLayout)
	}
	return headerRowStyle.Render(headerRowLayout)

}
