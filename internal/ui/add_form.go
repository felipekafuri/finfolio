package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type field struct {
	label       string
	value       string
	placeholder string
}

type AddFormModel struct {
	fields     []field
	focusIndex int
	submitted  bool
}

func NewAddForm() AddFormModel {
	return AddFormModel{
		fields: []field{
			{label: "Application Date", placeholder: "YYYY-MM-DD"},
			{label: "Value", placeholder: "1000.00"},
			{label: "Bank", placeholder: "Nubank"},
			{label: "Title", placeholder: "CDB 120% CDI"},
			{label: "Redemption Date", placeholder: "YYYY-MM-DD"},
		},
		focusIndex: 0,
	}
}

func (m AddFormModel) Init() tea.Cmd {
	return nil
}

// Update handles keyboard input and state changes
func (m AddFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			// Cancel form and quit
			return m, tea.Quit

		case "enter":
			// Move to next field or submit on last field
			if m.focusIndex == len(m.fields)-1 {
				m.submitted = true
				return m, tea.Quit
			}
			m.focusIndex++

		case "tab":
			// Move to next field (wrap to first)
			m.focusIndex++
			if m.focusIndex >= len(m.fields) {
				m.focusIndex = 0
			}

		case "shift+tab":
			// Move to previous field (wrap to last)
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.fields) - 1
			}

		case "backspace":
			// Delete last character from current field
			if len(m.fields[m.focusIndex].value) > 0 {
				m.fields[m.focusIndex].value = m.fields[m.focusIndex].value[:len(m.fields[m.focusIndex].value)-1]
			}

		default:
			// Add typed character to current field
			m.fields[m.focusIndex].value += msg.String()
		}
	}

	return m, nil
}

func (m AddFormModel) View() string {
	if m.submitted {
		return successStyle.Render("✓ Investment added successfully!\n")
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Add New Investment"))
	b.WriteString("\n\n")

	// Render each field
	for i, field := range m.fields {
		// Label
		label := field.label + ":"
		if i == m.focusIndex {
			label = focusedStyle.Render("▸ " + label)
		} else {
			label = blurredStyle.Render("  " + label)
		}
		b.WriteString(label + "\n")

		// Input value or placeholder
		value := field.value
		if value == "" {
			value = placeholderStyle.Render(field.placeholder)
		} else if i == m.focusIndex {
			value = focusedInputStyle.Render(value)
		} else {
			value = blurredInputStyle.Render(value)
		}
		b.WriteString("  " + value + "\n\n")
	}

	// Help text
	b.WriteString(helpStyle.Render("tab/shift+tab: navigate • enter: next/submit • esc: cancel"))

	return b.String()
}

func (m AddFormModel) GetInvestmentData() map[string]string {
	data := make(map[string]string)
	for _, field := range m.fields {
		data[field.label] = field.value
	}
	return data
}

func (m AddFormModel) IsSubmitted() bool {
	return m.submitted
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666"))

	focusedInputStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#7D56F4")).
				Padding(0, 1)

	blurredInputStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF"))

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666")).
				Italic(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666")).
			MarginTop(1)
)
