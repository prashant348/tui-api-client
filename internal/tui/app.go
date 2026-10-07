package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/prashant348/tui-api-client/internal/tui/components"
)

type Model struct {
	HTTPMethods  components.HTTPMethods
	URLInput     components.URLInput
	BodyInput    components.BodyInput
	HeadersInput components.HeadersInput
	SendButton   components.SendButton
	Footer       components.Footer
	FocusIndex   int
}

// dummy function to simulate http request
func sendHTTPRequest() {
	fmt.Println("\n\nSending HTTP request...")
}

func InitialModel() *Model {

	h := components.NewHTTPMethods()
	u := components.NewURLInput()
	h.Focus()
	s := components.NewSendButton(sendHTTPRequest)
	f := components.NewFooter()
	b := components.NewBodyInput()
	hi := components.NewHeadersInput()

	return &Model{
		HTTPMethods:  h,
		URLInput:     u,
		BodyInput:    b,
		HeadersInput: hi,
		SendButton:   s,
		Footer:       f,
		FocusIndex:   0,
	}
}

// use *Model to actually get the focusable components of real Model instead of irreleavent copy
func (m *Model) getFocusableComponents() map[int]FocusableComponent {
	return map[int]FocusableComponent{
		0: &m.HTTPMethods,
		1: &m.URLInput,
		2: &m.BodyInput,
		3: &m.HeadersInput,
		4: &m.SendButton,
	}
}

// using *Model, cuz, FocusNext is mutating FocusIndex State
func (m *Model) FocusNext() {
	components := m.getFocusableComponents()
	totalComponents := len(components)

	if currComp, exists := components[m.FocusIndex]; exists {
		currComp.Blur()
	}

	m.FocusIndex = (m.FocusIndex + 1) % totalComponents

	if m.FocusIndex == 2 && m.HTTPMethods.GetSelectedMethod() == "GET" {
		m.FocusIndex++
		m.FocusIndex++
	}

	if nextComp, exists := components[m.FocusIndex]; exists {
		nextComp.Focus()
	}

	// log.Printf("FocusNext(): current focus index: %v", m.FocusIndex)
}

// using Model instead of *Model, cause this func is only reading value not mutating their own state
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// using *Model to update actual Model struct intead of creating copy, updating, and returning it
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.FocusNext()
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.FocusIndex {
	case 0:
		m.HTTPMethods, cmd = m.HTTPMethods.Update(msg)
	case 1:
		m.URLInput, cmd = m.URLInput.Update(msg)
	case 2:
		m.BodyInput, cmd = m.BodyInput.Update(msg)
	case 3:
		m.HeadersInput, cmd = m.HeadersInput.Update(msg)
	case 4:
		m.SendButton, cmd = m.SendButton.Update(msg)
	}

	return m, cmd
}

// using Model instead of *Model, cause this func is only reading value not mutating its own state
func (m Model) View() string {

	selectedMethod := m.HTTPMethods.GetSelectedMethod()

	ui := m.URLInput.View()
	hm := m.HTTPMethods.View()
	sb := m.SendButton.View()
	fo := m.Footer.View()
	bi := m.BodyInput.View()
	hi := m.HeadersInput.View()

	// joining URLInput and SendButton Components horizontally
	ui_plus_sb := lipgloss.JoinHorizontal(
		lipgloss.Center,
		ui,
		sb,
	)

	var globalView string

	if selectedMethod != "GET" {
		globalView = lipgloss.JoinVertical(
			lipgloss.Left,
			hm,
			ui_plus_sb,
			bi,
			hi,
			fo,
		)
	} else {
		globalView = lipgloss.JoinVertical(
			lipgloss.Left,
			hm,
			ui_plus_sb,
			fo,
		)
	}

	return globalView
}
