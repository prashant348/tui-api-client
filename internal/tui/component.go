package tui

type FocusableComponent interface {
	Focus()
	Blur()
	IsFocused() bool
}
