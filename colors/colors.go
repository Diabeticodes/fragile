package colors

import "github.com/fatih/color"

// GetColorPrinter returns a color printer for the specified color name
func GetColorPrinter(colorName string) *color.Color {
	switch colorName {
	case "title":
		return color.New(color.FgRed)
	case "npc":
		return color.New(color.FgGreen)
	case "system":
		return color.New(color.FgYellow)
	case "bartender":
		return color.New(color.FgYellow)
	case "input":
		return color.New(color.FgCyan)
	default:
		return color.New()
	}
}

// Colorize applies color to text and returns the colored string
func Colorize(text, colorName string) string {
	return GetColorPrinter(colorName).Sprint(text)
}
