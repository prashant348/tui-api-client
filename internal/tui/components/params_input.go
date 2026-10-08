package components

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ParamsInput struct {
	Params         []ParamsRow
	isFocused      bool
	ActiveRowIndex int
}

func (p *ParamsInput) Focus() {
	p.isFocused = true
	p.Params[p.ActiveRowIndex].FocusRow() // for initial focus
}

func (p *ParamsInput) Blur() {
	p.isFocused = false
	// when global focus will shift from params input to send btn,
	// this Blur() will be called in FocusNext() func as currComp.Blur()
	// so with this func call, this logic will explicitily blur the last focused row
	if len(p.Params) > 0 && p.ActiveRowIndex >= 0 && p.ActiveRowIndex < len(p.Params) {
		p.Params[p.ActiveRowIndex].BlurRow()
	}
}

func (p ParamsInput) IsFocused() bool {
	return p.isFocused
}

func (p *ParamsInput) FocusPrevRow() {
	totalRows := len(p.Params)
	if totalRows == 0 || totalRows == 1 {
		return
	}

	if p.ActiveRowIndex >= 0 && p.ActiveRowIndex < totalRows {
		p.Params[p.ActiveRowIndex].BlurRow()
	}

	p.ActiveRowIndex = (p.ActiveRowIndex - 1 + totalRows) % totalRows
	p.Params[p.ActiveRowIndex].FocusRow()
}

func (p *ParamsInput) FocusNextRow() {
	totalRows := len(p.Params)
	if totalRows == 0 || totalRows == 1 {
		return
	}

	if p.ActiveRowIndex >= 0 && p.ActiveRowIndex < totalRows {
		p.Params[p.ActiveRowIndex].BlurRow()
	}

	p.ActiveRowIndex = (p.ActiveRowIndex + 1) % totalRows
	p.Params[p.ActiveRowIndex].FocusRow()
}

func NewParamsInput() ParamsInput {

	var params []ParamsRow

	for i := 0; i < 3; i++ {
		params = append(params, NewParamsRow())
	}

	return ParamsInput{
		Params:         params,
		isFocused:      false,
		ActiveRowIndex: 0,
	}
}

func (p ParamsInput) GetCurrRowIndex() int {
	return p.ActiveRowIndex
}

func (p ParamsInput) GetCurrRow() ParamsRow {
	return p.Params[p.ActiveRowIndex]
}

func (p *ParamsInput) AddRow() {
	newRow := NewParamsRow()
	p.Params = append(p.Params, newRow)
}

func (p *ParamsInput) RemoveRow() {
	// there should be at least one row
	if len(p.Params) == 1 {
		return
	}

	rowToRemove := p.ActiveRowIndex
	params := slices.Delete(p.Params, rowToRemove, rowToRemove+1)
	p.Params = params
	// after removing row, explicitly decrease active row index to avoid panic of index out of range
	if p.ActiveRowIndex >= len(p.Params) && len(p.Params) > 0 {
		p.ActiveRowIndex = len(p.Params) - 1
	}
	// after removing row, explicitly focus on the last/new row
	p.Params[p.ActiveRowIndex].FocusRow()
}

func (p *ParamsInput) Update(msg tea.Msg) (ParamsInput, tea.Cmd) {
	var cmd tea.Cmd

	if !p.IsFocused() {
		p.Params[p.ActiveRowIndex].BlurRow() // blur current focused row if component is not focused
		return *p, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+a":
			p.AddRow()
		case "ctrl+r":
			p.RemoveRow()
		case "down":
			p.FocusNextRow()
			return *p, cmd
		case "up":
			p.FocusPrevRow()
			return *p, cmd
		}
	}

	p.Params[p.ActiveRowIndex], cmd = p.Params[p.ActiveRowIndex].Update(msg)

	return *p, cmd

}

func (p ParamsInput) View() string {

	paramsInputFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	paramsInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	tagStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1)

	tag := "Params: "

	params := []string{tagStyle.Render(tag)}

	for i := range p.Params {
		params = append(params, p.Params[i].View())
	}

	var paramsLayout = lipgloss.JoinVertical(
		lipgloss.Left,
		params...,
	)

	if p.isFocused {
		return paramsInputFocusedStyle.Render(paramsLayout)
	}

	return paramsInputStyle.Render(paramsLayout)

}
