package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	keyvalueeditor "github.com/prashant348/tui-api-client/internal/tui/components/key_value_editor"
)

type Section interface {
	Focus()
	Blur()
	IsFocused() bool
	View() string
}

type Navigation struct {
	Tabs             []string
	SelectedTabIndex int
	isFocused        bool
	BodyInput        BodyInput
	Headers          Headers
	Params           Params
}

var tabs = []string{"Body", "Headers", "Params"}

var tabColorMapping = map[string]string{
	"Body":    "#33ff00",
	"Headers": "#ffd900",
	"Params":  "#00b7ff",
}

func (n *Navigation) Focus() {
	n.isFocused = true
	// instantly focus active section when the main component gets focused
	n.GetSections()[n.SelectedTabIndex].Focus()
}

func (n *Navigation) Blur() {
	n.isFocused = false
	// blur every section behind when main component gets blurred
	n.BodyInput.Blur()
	n.Headers.Blur()
	n.Params.Blur()
}

func (n *Navigation) IsFocused() bool {
	return n.isFocused
}

func (n *Navigation) GetSelectedTab() string {
	return n.Tabs[n.SelectedTabIndex]
}

func (n *Navigation) GetSections() map[int]Section {
	return map[int]Section{
		0: &n.BodyInput,
		1: &n.Headers,
		2: &n.Params,
	}
}

func NewNavigation() Navigation {

	b := NewBodyInput()

	headersConfig := keyvalueeditor.NewKeyValueEditorConfig(
		"Headers: ",
		"Key",
		"Value",
		3,
		50,
	)
	hs := NewHeaders(headersConfig)

	paramsConfig := keyvalueeditor.NewKeyValueEditorConfig(
		"Params: ",
		"Key",
		"Value",
		3,
		50,
	)
	ps := NewParams(paramsConfig)

	return Navigation{
		Tabs:             tabs,
		SelectedTabIndex: 0,
		isFocused:        false,
		BodyInput:        b,
		Headers:          hs,
		Params:           ps,
	}
}

func (n *Navigation) SwitchTab() {

	sections := n.GetSections()
	totalSections := len(sections)

	if currSec, ok := sections[n.SelectedTabIndex]; ok {
		currSec.Blur()
	}

	n.SelectedTabIndex = (n.SelectedTabIndex + 1) % totalSections

	if nextSec, ok := sections[n.SelectedTabIndex]; ok {
		nextSec.Focus()
	}

}

func (n *Navigation) Update(msg tea.Msg) (Navigation, tea.Cmd) {
	var cmd tea.Cmd

	if !n.IsFocused() {
		return *n, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "shift+tab":
			n.SwitchTab()
		}
	}
	switch n.SelectedTabIndex {
	case 0:
		n.BodyInput, cmd = n.BodyInput.Update(msg)
	case 1:
		n.Headers, cmd = n.Headers.Update(msg)
	case 2:
		n.Params, cmd = n.Params.Update(msg)
	}

	return *n, cmd
}

func (n Navigation) View() string {

	currTab := n.GetSelectedTab()

	tabStyle := lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder()).
		Padding(0, 1)

	selectedTabStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color(tabColorMapping[currTab])).
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
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#4B5563")).
		Padding(0, 1)

	tag := "Tabs: "
	var tabs []string

	for i, method := range n.Tabs {
		if i == n.SelectedTabIndex {
			tabs = append(tabs, selectedTabStyle.Render(method))
		} else {
			tabs = append(tabs, tabStyle.Render(method))
		}
	}

	styledTagView := tagStyle.Render(tag)
	tabsLayout := lipgloss.JoinHorizontal(lipgloss.Center, tabs...)

	tabsBar := lipgloss.JoinHorizontal(
		lipgloss.Center,
		styledTagView,
		tabsLayout,
	)

	var comp string = n.GetSections()[n.SelectedTabIndex].View()

	finalLayout := lipgloss.JoinVertical(
		lipgloss.Left,
		tabsBar,
		comp,
	)

	if n.isFocused {
		return ComponentFocusedStyle.Render(finalLayout)
	}

	return ComponentStyle.Render(finalLayout)
}
