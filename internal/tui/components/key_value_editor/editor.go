package keyvalueeditor

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyValueEditor struct {
	Rows           []KeyValueRow
	isFocused      bool
	ActiveRowIndex int
	Config         KeyValueEditorConfig
}

func (e *KeyValueEditor) Focus() {
	e.isFocused = true
	e.Rows[e.ActiveRowIndex].FocusRow()
}

func (e *KeyValueEditor) Blur() {
	e.isFocused = false
	if len(e.Rows) > 0 && e.ActiveRowIndex >= 0 && e.ActiveRowIndex < len(e.Rows) {
		e.Rows[e.ActiveRowIndex].BlurRow()
	}
}

func (e KeyValueEditor) IsFocused() bool {
	return e.isFocused
}

func (e *KeyValueEditor) ShiftRowFocus(key string) {
	totalRows := len(e.Rows)
	// there should be atleast 2 rows to shift focus across them
	if totalRows <= 1 {
		return
	}

	if e.ActiveRowIndex >= 0 && e.ActiveRowIndex < totalRows {
		e.Rows[e.ActiveRowIndex].BlurRow()
	}

	switch key {
	case "up":
		e.ActiveRowIndex = (e.ActiveRowIndex - 1 + totalRows) % totalRows
	case "down":
		e.ActiveRowIndex = (e.ActiveRowIndex + 1) % totalRows
	}

	e.Rows[e.ActiveRowIndex].FocusRow()
}

func NewKeyValueEditor(config KeyValueEditorConfig) KeyValueEditor {

	var rows []KeyValueRow

	for i := 0; i < config.InitialRows; i++ {
		rows = append(rows, NewKeyValueRow(config))
	}

	return KeyValueEditor{
		Rows:           rows,
		isFocused:      false,
		ActiveRowIndex: 0,
		Config:         config,
	}
}

func (e *KeyValueEditor) AddRow() {
	newRow := NewKeyValueRow(e.Config)
	e.Rows = append(e.Rows, newRow)
}

func (e *KeyValueEditor) RemoveRow() {
	// there should be atleast 1 row
	if len(e.Rows) <= 1 {
		return
	}
	// remove the active row
	rowToRemove := e.ActiveRowIndex
	rows := slices.Delete(e.Rows, rowToRemove, rowToRemove+1)
	e.Rows = rows // assign updated rows

	// reduce ActiveRowIndex by 1 to point to the last/new row to avoid index out of bounds error
	if e.ActiveRowIndex >= len(e.Rows) && len(e.Rows) > 0 {
		e.ActiveRowIndex = len(e.Rows) - 1
	}
	// and then focus the last/new row
	e.Rows[e.ActiveRowIndex].FocusRow()
}

func (e *KeyValueEditor) Update(msg tea.Msg) (KeyValueEditor, tea.Cmd) {
	var cmd tea.Cmd

	if !e.IsFocused() {
		e.Rows[e.ActiveRowIndex].BlurRow() // explicitly blur the active row when editor comp is not focused
		return *e, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+a":
			e.AddRow()
		case "ctrl+r":
			e.RemoveRow()
		case "up", "down":
			e.ShiftRowFocus(msg.String())
			return *e, cmd
		}
	}

	e.Rows[e.ActiveRowIndex], cmd = e.Rows[e.ActiveRowIndex].Update(msg)

	return *e, cmd
}

func (e KeyValueEditor) View() string {
	editorFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	editorNormalStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1)

	title := e.Config.Title

	rows := []string{titleStyle.Render(title)}

	for i := range e.Rows {
		rows = append(rows, e.Rows[i].View())
	}

	var editorLayout = lipgloss.JoinVertical(
		lipgloss.Left,
		rows...,
	)

	if e.isFocused {
		return editorFocusedStyle.Render(editorLayout)
	}

	return editorNormalStyle.Render(editorLayout)
}
