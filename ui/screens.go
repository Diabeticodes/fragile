package ui

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"

	"github.com/michaelmastrucci/text-rpg/colors"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ShowStartScreen displays the game's start screen with ASCII art and menu options
func ShowStartScreen() string {
	ClearScreen()

	// ASCII Art for "FRAGILE"
	asciiArt := `
	FFFFFFF  RRRRRR       A       GGGGG   IIIII   L       EEEEE
	F        R     R     A A      G         I     L       E
	F        R     R    A   A     G  GGG    I     L       E
	FFFR     R  RRR    A     A    G    G    I     L       EEE
	F        R  R     A       A   G    G    I     L       E
	F        R   R   AAAAAAAAAAA   GGGG   IIIII   LLLLL   EEEEE
`

	// Display ASCII art in title color
	TypeWriter(asciiArt, 0, colors.GetColorPrinter("title"))
	
	// Menu options
	menuOptions := []string{
		"> Start",
		"> Continue",
		"> Options",
	}

	// Display menu options
	for _, option := range menuOptions {
		TypeWriter(option, 20, colors.GetColorPrinter("bartender"))
	}

	// Get user input
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if len(input) > 0 {
			switch input[0] {
			case 's', 'S':
				return "start"
			case 'l', 'L':
				return "load"
			case 'o', 'O':
				return "options"
			default:
				TypeWriter("Invalid option. Please try again.", 20, colors.GetColorPrinter("system"))
				continue
			}
		}
	}
}
