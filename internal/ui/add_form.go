package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type field struct {
	label       string
	value       string
	placeholder string
	icon        string
}

type AddFormModel struct {
	fields     []field
	focusIndex int
	submitted  bool
	width      int
	height     int
}

func NewAddForm() AddFormModel {
	return AddFormModel{
		fields: []field{
			{label: "Application Date", placeholder: "YYYY/MM/DD", icon: "📅"},
			{label: "Value", placeholder: "1000.00", icon: "💰"},
			{label: "Bank", placeholder: "Nubank", icon: "🏦"},
			{label: "Title", placeholder: "CDB 120% CDI", icon: "📊"},
			{label: "Redemption Date", placeholder: "YYYY/MM/DD", icon: "🎯"},
		},
		focusIndex: 0,
		width:      80,
		height:     25,
	}
}

func (m AddFormModel) Init() tea.Cmd {
	return nil
}

// Update handles keyboard input and state changes
func (m AddFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// If already submitted, any key quits
		if m.submitted {
			return m, tea.Quit
		}

		switch msg.String() {
		case "ctrl+c", "esc":
			// Cancel form and quit
			return m, tea.Quit

		case "enter":
			// Move to next field or submit on last field
			if m.focusIndex == len(m.fields)-1 {
				m.submitted = true
				return m, nil // Don't quit immediately, show success screen first
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
			// Add typed character to current field (filter out non-printable)
			if len(msg.String()) == 1 {
				m.fields[m.focusIndex].value += msg.String()
			}
		}
	}

	return m, nil
}

func (m AddFormModel) View() string {
	if m.submitted {
		return m.renderSuccessView()
	}

	return m.renderFormView()
}

func (m AddFormModel) renderFormView() string {
	var b strings.Builder

	// Header with gradient effect
	header := headerStyle.Render("💼 FINFOLIO")
	subtitle := subtitleStyle.Render("Add New Investment")

	b.WriteString("\n")
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, header))
	b.WriteString("\n")
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, subtitle))
	b.WriteString("\n\n")

	// Progress indicator
	progress := m.renderProgressBar()
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, progress))
	b.WriteString("\n\n")

	// Form fields container
	var formContent strings.Builder

	for i, field := range m.fields {
		formContent.WriteString(m.renderField(i, field))
		if i < len(m.fields)-1 {
			formContent.WriteString("\n")
		}
	}

	// Wrap form in a bordered box
	formBox := formBoxStyle.Render(formContent.String())
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, formBox))
	b.WriteString("\n\n")

	// Help text with icons
	help := m.renderHelp()
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, help))
	b.WriteString("\n")

	return b.String()
}

func (m AddFormModel) renderField(index int, field field) string {
	isFocused := index == m.focusIndex

	// Label with icon
	labelText := fmt.Sprintf("%s  %s", field.icon, field.label)
	var label string
	if isFocused {
		label = focusedLabelStyle.Render("▸ " + labelText)
	} else {
		label = blurredLabelStyle.Render("  " + labelText)
	}

	// Input box
	value := field.value
	var inputBox string

	if value == "" {
		placeholder := placeholderStyle.Render(field.placeholder)
		if isFocused {
			inputBox = focusedInputBoxStyle.Render(placeholder + cursorStyle.Render("█"))
		} else {
			inputBox = blurredInputBoxStyle.Render(placeholder)
		}
	} else {
		if isFocused {
			inputBox = focusedInputBoxStyle.Render(value + cursorStyle.Render("█"))
		} else {
			inputBox = blurredInputBoxStyle.Render(value)
		}
	}

	return label + "\n" + inputBox
}

func (m AddFormModel) renderProgressBar() string {
	total := len(m.fields)
	current := m.focusIndex + 1

	var bar strings.Builder
	bar.WriteString(progressTextStyle.Render(fmt.Sprintf("Step %d of %d  ", current, total)))

	for i := 0; i < total; i++ {
		if i < current {
			bar.WriteString(progressFilledStyle.Render("●"))
		} else if i == current {
			bar.WriteString(progressCurrentStyle.Render("●"))
		} else {
			bar.WriteString(progressEmptyStyle.Render("○"))
		}
		if i < total-1 {
			bar.WriteString(" ")
		}
	}

	return bar.String()
}

func (m AddFormModel) renderHelp() string {
	keys := []string{
		"↹ tab/shift+tab: navigate",
		"↵ enter: next/submit",
		"⎋ esc: cancel",
	}

	return helpBoxStyle.Render(strings.Join(keys, "  •  "))
}

func (m AddFormModel) renderSuccessView() string {
	var b strings.Builder

	// Clear previous content
	b.WriteString("\033[2J\033[H") // ANSI clear screen and move to top

	// Success animation/banner
	banner := successBannerStyle.Render(`
    ✓ SUCCESS!
    ═══════════════════════════════════

    Investment Added Successfully

    Your investment has been recorded
    and is now part of your portfolio.
    `)

	b.WriteString("\n\n\n")
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, banner))
	b.WriteString("\n\n")

	// Build cards as separate elements
	var row1Cards []string
	var row2Cards []string

	for i, field := range m.fields {
		if field.value != "" {
			card := dataCardStyle.Render(
				fmt.Sprintf("%s  %s\n%s",
					field.icon,
					dimTextStyle.Render(field.label),
					valueTextStyle.Render(field.value),
				),
			)

			if i < 3 {
				row1Cards = append(row1Cards, card)
			} else {
				row2Cards = append(row2Cards, card)
			}
		}
	}

	// Join cards horizontally and center each row
	if len(row1Cards) > 0 {
		row1 := lipgloss.JoinHorizontal(lipgloss.Top, row1Cards...)
		b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, row1))
		b.WriteString("\n")
	}
	if len(row2Cards) > 0 {
		row2 := lipgloss.JoinHorizontal(lipgloss.Top, row2Cards...)
		b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top, row2))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Top,
		helpStyle.Render("Press any key to exit...")))
	b.WriteString("\n")

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

// Color palette
var (
	primaryColor   = lipgloss.Color("#7D56F4")
	secondaryColor = lipgloss.Color("#FF6B9D")
	successColor   = lipgloss.Color("#04B575")
	mutedColor     = lipgloss.Color("#626262")
	bgColor        = lipgloss.Color("#1A1A2E")
	fgColor        = lipgloss.Color("#EAEAEA")
	accentColor    = lipgloss.Color("#FFD93D")
)

// Styles
var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.AdaptiveColor{Light: "#7D56F4", Dark: "#7D56F4"}).
			Padding(0, 2).
			MarginBottom(0)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)

	formBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			Width(60)

	focusedLabelStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				MarginBottom(1)

	blurredLabelStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				MarginBottom(1)

	focusedInputBoxStyle = lipgloss.NewStyle().
				Foreground(fgColor).
				Background(primaryColor).
				Padding(0, 1).
				Width(54).
				Bold(true)

	blurredInputBoxStyle = lipgloss.NewStyle().
				Foreground(fgColor).
				Border(lipgloss.NormalBorder()).
				BorderForeground(mutedColor).
				Padding(0, 1).
				Width(54)

	placeholderStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Italic(true)

	cursorStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Blink(true)

	progressTextStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				MarginRight(1)

	progressFilledStyle = lipgloss.NewStyle().
				Foreground(successColor)

	progressCurrentStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true)

	progressEmptyStyle = lipgloss.NewStyle().
				Foreground(mutedColor)

	helpBoxStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(mutedColor).
			Padding(0, 2)

	helpStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	successBannerStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true).
				Border(lipgloss.DoubleBorder()).
				BorderForeground(successColor).
				Padding(1, 4).
				Align(lipgloss.Center)

	dataCardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			MarginRight(1)

	dimTextStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginBottom(1)

	valueTextStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)
)
