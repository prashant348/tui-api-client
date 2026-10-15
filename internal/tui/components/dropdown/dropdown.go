package dropdown

import (
	"errors"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dropdown struct {
	Options             []string
	SelectedOptionIndex int
	CursorIndex         int
	isFocused           bool
	isOpened            bool
}

func NewDropdown(
	options []string,
	defaultOptionIndex int,
) (Dropdown, error) {

	options = slices.DeleteFunc(options, func(option string) bool {
		return strings.TrimSpace(option) == ""
	})

	if len(options) == 0 {
		return Dropdown{}, errors.New("dropdown requires at least one non-empty option")
	}

	if defaultOptionIndex < 0 || defaultOptionIndex >= len(options) {
		defaultOptionIndex = 0
	}

	return Dropdown{
		Options:             options,
		SelectedOptionIndex: defaultOptionIndex,
		CursorIndex:         defaultOptionIndex,
		isFocused:           false,
		isOpened:            false,
	}, nil
}

func (d *Dropdown) Focus() {
	d.isFocused = true
}

func (d *Dropdown) Blur() {
	d.isFocused = false
	// dont forget to close the dropdown menu when it is not focused
	d.Close()
}

func (d Dropdown) IsFocused() bool {
	return d.isFocused
}

func (d *Dropdown) Open() {
	d.isOpened = true
}

func (d *Dropdown) Close() {
	d.isOpened = false
}

func (d Dropdown) IsOpen() bool {
	return d.isOpened
}

func (d Dropdown) GetSelectedItem() string {
	return d.Options[d.SelectedOptionIndex]
}

func (d *Dropdown) MoveCursor(key string) {

	totalOptions := len(d.Options)

	switch key {
	case "up":
		d.CursorIndex = (d.CursorIndex - 1 + totalOptions) % totalOptions
	case "down":
		d.CursorIndex = (d.CursorIndex + 1) % totalOptions
	}

	// d.SelectedOptionIndex = d.CursorIndex
}

func (d *Dropdown) SelectOption() {
	d.SelectedOptionIndex = d.CursorIndex
}

func (d *Dropdown) Update(msg tea.Msg) (Dropdown, tea.Cmd) {
	var cmd tea.Cmd

	if !d.IsFocused() {
		return *d, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if d.IsOpen() {
				d.SelectOption()
			} else {
				d.Open()
			}
		case "esc":
			d.Close()
		case "up", "down":
			if d.IsOpen() {
				d.MoveCursor(msg.String())
			}
		}
	}

	return *d, cmd
}

func (d Dropdown) View() string {
	button := d.DropdownButton()

	if !d.IsOpen() {
		return button
	}

	menu := d.DropdownMenu()

	view := lipgloss.JoinVertical(
		lipgloss.Center,
		button,
		menu,
	)

	return view

}

func (d Dropdown) DropdownButton() string {
	totalWidth := 14

	dropdownButtonStyle := lipgloss.NewStyle().
		// Border(lipgloss.HiddenBorder()).
		Background(lipgloss.Color("#6f42c1")).
		Padding(0, 1).Width(totalWidth)

	dropdownButtonFocusedStyle := lipgloss.NewStyle().
		// Border(lipgloss.HiddenBorder()).
		Background(lipgloss.Color("#5f16e6")).
		Padding(0, 1).Width(totalWidth)

	downsideTriangle := "▼"
	// leftsideTriangle := "◀"
	upsideTriangle := "▲"
	// leftTriangle := "◀"
	// rightTriangle := "▶"
	selectedItem := d.GetSelectedItem()

	spaceCount := totalWidth - len(selectedItem) - len(downsideTriangle)

	if spaceCount < 1 {
		spaceCount = 1
	}

	dynamicSpace := lipgloss.PlaceHorizontal(spaceCount, lipgloss.Left, "")

	var triangle string

	if d.IsOpen() {
		triangle = upsideTriangle
	} else {
		triangle = downsideTriangle
	}

	dropdownButtonLayout := lipgloss.JoinHorizontal(
		lipgloss.Center,
		selectedItem,
		dynamicSpace,
		triangle,
	)

	if d.IsFocused() {
		return dropdownButtonFocusedStyle.Render(dropdownButtonLayout)
	}

	return dropdownButtonStyle.Render(dropdownButtonLayout)
}

func (d Dropdown) DropdownMenu() string {
	// totalWidth := 14

	dropdownMenuStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#5900ff")).
		Padding(0, 1)

	dropdownSelectedOptionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5900ff")).
		Bold(true).Width(8)

	dropdownOptionStyle := lipgloss.NewStyle().
		// Foreground(lipgloss.Color("#fff")).
		Padding(0, 0).Bold(false)

	cursorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Bold(true)

	rightArrow := "▶"
	cursor := cursorStyle.Render(rightArrow)

	var options []string

	for i, option := range d.Options {

		if i == d.SelectedOptionIndex {
			if d.CursorIndex == d.SelectedOptionIndex {
				options = append(options, cursor+" "+dropdownSelectedOptionStyle.Render(option))
			} else {
				options = append(options, "  "+dropdownSelectedOptionStyle.Render(option))
			}
		} else {
			if i == d.CursorIndex {
				options = append(options, cursor+" "+dropdownOptionStyle.Render(option))
			} else {
				options = append(options, "  "+dropdownOptionStyle.Render(option))
			}
		}
	}

	menuLayout := lipgloss.JoinVertical(
		lipgloss.Left,
		options...,
	)

	return dropdownMenuStyle.Render(menuLayout)

}
