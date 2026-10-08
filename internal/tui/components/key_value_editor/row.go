package keyvalueeditor

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyValueRow struct {
	KeyInput       textinput.Model
	ValueInput     textinput.Model
	isFocused      bool
	isKeyFocused   bool
	isValueFocused bool
}

func (r *KeyValueRow) FocusRow() {
	r.isFocused = true
	r.FocusRowKey()  // default focus is key input
	r.BlurRowValue() // and explicitily blur value input when key input is focused
}

func (r *KeyValueRow) BlurRow() {
	r.isFocused = false
	r.BlurRowKey()
	r.BlurRowValue()
}

func (r KeyValueRow) IsFocusedRow() bool {
	return r.isFocused
}

func (r *KeyValueRow) FocusRowKey() {
	r.isKeyFocused = true
	r.KeyInput.Focus()
}

func (r *KeyValueRow) FocusRowValue() {
	r.isValueFocused = true
	r.ValueInput.Focus()
}

func (r *KeyValueRow) BlurRowKey() {
	r.isKeyFocused = false
	r.KeyInput.Blur()
}

func (r *KeyValueRow) BlurRowValue() {
	r.isValueFocused = false
	r.ValueInput.Blur()
}

func NewKeyValueRow(config KeyValueEditorConfig) KeyValueRow {

	ki := textinput.New()

	ki.Placeholder = config.KeyPlaceholder
	ki.CharLimit = 256
	ki.Width = config.Width / 2

	vi := textinput.New()

	vi.Placeholder = config.ValuePlaceholder
	vi.CharLimit = 256
	vi.Width = config.Width / 2

	keyValueRow := KeyValueRow{
		KeyInput:   ki,
		ValueInput: vi,
		isFocused:  false,
	}

	return keyValueRow
}

func (r *KeyValueRow) Update(msg tea.Msg) (KeyValueRow, tea.Cmd) {
	var cmd tea.Cmd
	if !r.IsFocusedRow() {
		return *r, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if r.isKeyFocused {
				r.BlurRowKey()
				r.FocusRowValue()
			} else if r.isValueFocused {
				r.BlurRowValue()
				r.FocusRowKey()
			}
		}
	}

	if r.isKeyFocused {
		r.KeyInput, cmd = r.KeyInput.Update(msg)
	} else if r.isValueFocused {
		r.ValueInput, cmd = r.ValueInput.Update(msg)
	}

	return *r, cmd
}

func (r KeyValueRow) View() string {

	rowStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	rowFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	rowLayout := lipgloss.JoinHorizontal(
		lipgloss.Center,
		r.KeyInput.View(),
		r.ValueInput.View(),
	)

	if r.IsFocusedRow() {
		return rowFocusedStyle.Render(rowLayout)
	}
	return rowStyle.Render(rowLayout)

}
