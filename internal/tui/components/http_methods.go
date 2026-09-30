package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HTTPMethods struct {
	Methods       []string
	SelectedIndex int
	isFocused     bool
}

var httpMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

var methodColorMapping = map[string]string{
	"GET":    "#33ff00",
	"POST":   "#ffd900",
	"PUT":    "#00b7ff",
	"DELETE": "#ff0202",
	"PATCH":  "#b700ff",
}

func (h *HTTPMethods) Focus() {
	h.isFocused = true
}

func (h *HTTPMethods) Blur() {
	h.isFocused = false
}

func (h HTTPMethods) IsFocused() bool {
	return h.isFocused
}

func (h *HTTPMethods) getSelectedMethod() string {
	return httpMethods[h.SelectedIndex]
}

func NewHTTPMethods() HTTPMethods {
	return HTTPMethods{
		Methods:       httpMethods,
		SelectedIndex: 0,
		isFocused:     false,
	}
}

func (h *HTTPMethods) Update(msg tea.Msg) (HTTPMethods, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right", "k":
			h.SelectedIndex = (h.SelectedIndex + 1) % len(h.Methods)
		case "left", "j":
			h.SelectedIndex = (h.SelectedIndex - 1 + len(h.Methods)) % len(h.Methods)
		}
	}

	return *h, cmd
}

func (h HTTPMethods) View() string {

	currMethod := h.getSelectedMethod()

	methodStyle := lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Padding(0, 1)

	selectedMethodStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color(methodColorMapping[currMethod])).
		Padding(0, 1).
		Bold(true).Foreground(lipgloss.Color("#fff"))

	tagStyle := lipgloss.NewStyle().
		Bold(true).
		// Foreground(lipgloss.Color("#fff")).
		Padding(0, 1)

	ComponentFocusedStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#6f42c1")).
		Padding(0, 1)

	ComponentStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	tag := "Methods: "
	var methods []string

	for i, method := range h.Methods {
		if i == h.SelectedIndex {
			methods = append(methods, selectedMethodStyle.Render(method))
		} else {
			methods = append(methods, methodStyle.Render(method))
		}
	}

	styledTagView := tagStyle.Render(tag)
	methodsView := lipgloss.JoinHorizontal(lipgloss.Center, methods...)

	finalView := lipgloss.JoinHorizontal(
		lipgloss.Center,
		styledTagView,
		methodsView,
	)

	if h.isFocused {
		return ComponentFocusedStyle.Render(finalView)
	}

	return ComponentStyle.Render(finalView)
}
