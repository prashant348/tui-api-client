package components

import (
	tea "github.com/charmbracelet/bubbletea"
	keyvalueeditor "github.com/prashant348/tui-api-client/internal/tui/components/key_value_editor"
)

type Headers struct {
	Editor keyvalueeditor.KeyValueEditor
}

func NewHeaders(config keyvalueeditor.KeyValueEditorConfig) Headers {
	return Headers{Editor: keyvalueeditor.NewKeyValueEditor(config)}
}

func (h *Headers) Focus() {
	h.Editor.Focus()
}

func (h *Headers) Blur() {
	h.Editor.Blur()
}

func (h Headers) IsFocused() bool {
	return h.Editor.IsFocused()
}

func (h *Headers) Update(msg tea.Msg) (Headers, tea.Cmd) {
	var cmd tea.Cmd
	h.Editor, cmd = h.Editor.Update(msg)
	return *h, cmd
}

func (h Headers) View() string {
	return h.Editor.View()
}
