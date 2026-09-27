package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// list of methods
var Methods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

// model struct in which application state is stored
type Model struct {
	UrlInput       textinput.Model // URL daalne ke liye input box
	SelectedMethod int             // Index of selected method (0 = GET, 1 = POST, etc.)
	FocusIndex     int             // 0 = Method, 1 = URL Input
	ResponseBody   string          // API Response yahan dikhega
	StatusCode     int             // e.g. 200, 404
	IsLoading      bool            // Spinner status
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	return Model{
		UrlInput:       ti,
		SelectedMethod: 0,
		FocusIndex:     0,
		ResponseBody:   "Press [Enter] or [Ctrl+S] to send request...\nPress [Tab] to switch focus.",
		StatusCode:     0,
		IsLoading:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink // start cursor blink animation
}
