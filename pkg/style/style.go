package style

import "github.com/charmbracelet/lipgloss"

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	HelpStyle  = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})
	ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	Bold       = lipgloss.NewStyle().Bold(true)
	ChangeFrom = lipgloss.NewStyle().Background(lipgloss.Color("88")).Foreground(lipgloss.Color("#ffffff"))
	ChangeTo   = lipgloss.NewStyle().Background(lipgloss.Color("28")).Foreground(lipgloss.Color("#ffffff"))
	Base       = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("238")).
			Align(lipgloss.Left)
	ActiveStyleBase = lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("248")).
			Align(lipgloss.Left)
	CostStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	SavingStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	InputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	ContinueStyle = lipgloss.NewStyle().Foreground(darkGray)
	SvcDisable    = lipgloss.NewStyle().Background(lipgloss.Color("#222222"))
	SvcEnable     = lipgloss.NewStyle().Background(lipgloss.Color("#aa2222"))
)