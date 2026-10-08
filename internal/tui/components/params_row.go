package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ParamsRow struct {
	KeyInput       textinput.Model
	ValueInput     textinput.Model
	isFocused      bool
	isKeyFocused   bool
	isValueFocused bool
}

func (p *ParamsRow) FocusRow() {
	p.isFocused = true
	p.FocusRowKey()  // default focus is key input
	p.BlurRowValue() // and explicitily blur value input when key input is focused
}

func (p *ParamsRow) BlurRow() {
	p.isFocused = false
	p.BlurRowKey()
	p.BlurRowValue()
}

func (p ParamsRow) IsFocusedRow() bool {
	return p.isFocused
}

func (p *ParamsRow) FocusRowKey() {
	p.isKeyFocused = true
	p.KeyInput.Focus()
}

func (p *ParamsRow) FocusRowValue() {
	p.isValueFocused = true
	p.ValueInput.Focus()
}

func (p *ParamsRow) BlurRowKey() {
	p.isKeyFocused = false
	p.KeyInput.Blur()
}

func (p *ParamsRow) BlurRowValue() {
	p.isValueFocused = false
	p.ValueInput.Blur()
}

func NewParamsRow() ParamsRow {

	ki := textinput.New()

	ki.Placeholder = "Key"
	ki.CharLimit = 256
	ki.Width = 24

	vi := textinput.New()

	vi.Placeholder = "Value"
	vi.CharLimit = 256
	vi.Width = 24

	paramsRow := ParamsRow{
		KeyInput:   ki,
		ValueInput: vi,
		isFocused:  false,
	}

	return paramsRow
}

func (p *ParamsRow) Update(msg tea.Msg) (ParamsRow, tea.Cmd) {
	var cmd tea.Cmd
	if !p.IsFocusedRow() {
		return *p, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if p.isKeyFocused {
				p.BlurRowKey()
				p.FocusRowValue()
			} else if p.isValueFocused {
				p.BlurRowValue()
				p.FocusRowKey()
			}
			// fmt.Printf("key focus: %v, value focus: %v", p.isKeyFocused, p.isValueFocused)
		}
	}

	if p.isKeyFocused {
		p.KeyInput, cmd = p.KeyInput.Update(msg)
	} else if p.isValueFocused {
		p.ValueInput, cmd = p.ValueInput.Update(msg)
	}

	return *p, cmd
}

func (p ParamsRow) View() string {

	paramRowStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	paramRowFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	paramRowLayout := lipgloss.JoinHorizontal(
		lipgloss.Center,
		p.KeyInput.View(),
		p.ValueInput.View(),
	)

	if p.IsFocusedRow() {
		return paramRowFocusedStyle.Render(paramRowLayout)
	}
	return paramRowStyle.Render(paramRowLayout)

}
