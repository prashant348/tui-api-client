package components

import (
	tea "github.com/charmbracelet/bubbletea"
	keyvalueeditor "github.com/prashant348/tui-api-client/internal/tui/components/key_value_editor"
)

type Params struct {
	Editor keyvalueeditor.KeyValueEditor
}

func NewParams(config keyvalueeditor.KeyValueEditorConfig) Params {
	return Params{Editor: keyvalueeditor.NewKeyValueEditor(config)}
}

func (p *Params) Focus() {
	p.Editor.Focus()
}

func (p *Params) Blur() {
	p.Editor.Blur()
}

func (p Params) IsFocused() bool {
	return p.Editor.IsFocused()
}

func (p *Params) Update(msg tea.Msg) (Params, tea.Cmd) {
	var cmd tea.Cmd
	p.Editor, cmd = p.Editor.Update(msg)
	return *p, cmd
}

func (p Params) View() string {
	return p.Editor.View()
}
