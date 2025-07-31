package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/michaelmastrucci/text-rpg/colors"
)

// TypeWriter prints text with a typewriter effect using the provided color
func TypeWriter(text string, speed time.Duration, c *color.Color) {
	c.EnableColor()
	for _, char := range text {
		c.Print(string(char))
		time.Sleep(speed * time.Millisecond)
	}
	fmt.Println()
}

// TypeWriterWithHighlight prints text with a typewriter effect, highlighting specific parts
func TypeWriterWithHighlight(text string, highlightText string, highlightColor string, speed time.Duration, baseColor *color.Color) {
	parts := strings.SplitN(text, highlightText, 2)
	if len(parts) == 2 {
		TypeWriter(parts[0], speed, baseColor)
		TypeWriter(highlightText, speed, colors.GetColorPrinter(highlightColor))
		TypeWriter(parts[1], speed, baseColor)
	} else {
		TypeWriter(text, speed, baseColor)
	}
}
