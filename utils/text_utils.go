package utils

import (
	"strings"
)

// WrapTextInBorder wraps the given text in a border with the specified width
func WrapTextInBorder(text string, width int) string {
	if width < 20 {
		width = 60 // Minimum width to ensure readability
	}

	// Split text into lines first
	inputLines := strings.Split(text, "\n")
	var lines []string

	// Process each line to handle word wrapping
	for _, line := range inputLines {
		words := strings.Fields(line)
		currentLine := ""

		for _, word := range words {
			// If adding this word would exceed the width, add current line and start new line
			if len(currentLine) > 0 && len(currentLine)+len(word)+1 > width-4 {
				lines = append(lines, currentLine)
				currentLine = word
			} else {
				if currentLine != "" {
					currentLine += " "
				}
				currentLine += word
			}
		}

		// Add the last line if not empty
		if currentLine != "" {
			lines = append(lines, currentLine)
		}

		// Add an empty line after each paragraph (except the last one)
		if line == "" && len(lines) > 0 && lines[len(lines)-1] != "" {
			lines = append(lines, "")
		}
	}

	// Create the top and bottom borders
	topBorder := "╭" + strings.Repeat("─", width-2) + "╮\n"
	bottomBorder := "╰" + strings.Repeat("─", width-2) + "╯"

	// Build the content with left and right borders
	var content strings.Builder
	for _, line := range lines {
		// Ensure the line doesn't exceed the width
		if len(line) > width-4 {
			line = line[:width-4]
		}
		// Add left border, line, and right border with padding
		content.WriteString("│ " + line + strings.Repeat(" ", width-len(line)-3) + "│\n")
	}

	// Combine everything
	return topBorder + content.String() + bottomBorder
}
