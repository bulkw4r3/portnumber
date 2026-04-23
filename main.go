package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginLeft(2)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5555"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))
)

type model struct {
	textInput textinput.Model
	port      int
	ip        string
	err       string
	quitting  bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "1-48"
	ti.Focus()
	ti.CharLimit = 2
	ti.Width = 10
	ti.Prompt = ""

	return model{
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			if m.err == "" && m.ip != "" {
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.validate()

	return m, cmd
}

func (m *model) validate() {
	val := strings.TrimSpace(m.textInput.Value())
	if val == "" {
		m.err = ""
		m.ip = ""
		return
	}

	port, err := strconv.Atoi(val)
	if err != nil {
		m.err = "Please enter a number"
		m.ip = ""
		return
	}

	if port < 1 || port > 48 {
		m.err = "Port must be between 1 and 48"
		m.ip = ""
		return
	}

	m.port = port
	m.ip = fmt.Sprintf("172.16.76.%d", 48+port)
	m.err = ""
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString(titleStyle.Render("Portnumber to IP Converter"))
	b.WriteString("\n\n")

	b.WriteString(labelStyle.Render("Port: "))
	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	if m.err != "" {
		b.WriteString(errorStyle.Render("✗ " + m.err))
		b.WriteString("\n")
	}

	if m.ip != "" {
		b.WriteString(labelStyle.Render("IP:   "))
		b.WriteString(resultStyle.Render(m.ip))
		b.WriteString("\n\n")
		b.WriteString(helpStyle.Render("Press Enter to confirm"))
	} else {
		b.WriteString(helpStyle.Render("Enter a port number (1-48)"))
	}

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("esc/ctrl+c to quit"))
	b.WriteString("\n")

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if mm, ok := m.(model); ok && mm.quitting && mm.ip != "" {
		fmt.Printf("Portnumber %d -> %s\n", mm.port, mm.ip)
	}
}
