package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
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
	spinner   spinner.Model
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

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))

	return model{
		textInput: ti,
		spinner:   s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

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

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.textInput, _ = m.textInput.Update(msg)
	m.validate()

	return m, tea.Batch(cmds...)
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
	b.WriteString(m.spinner.View())
	b.WriteString(" ")
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
