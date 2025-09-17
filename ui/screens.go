package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/eiannone/keyboard"
	"github.com/text-rpg/colors"
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
		menuOptions := []string{"Start", "Continue", "Options"}
		selected := 0

		asciiArt := "FFFFFFF  RRRRRR       A       GGGGG   IIIII   L       EEEEE\n" +
			"F        R     R     A A      G         I     L       E\n" +
			"F        R     R    A   A     G  GGG    I     L       E\n" +
			"FFFR     R  RRR    A     A    G    G    I     L       EEE\n" +
			"F        R  R     A       A   G    G    I     L       E\n" +
			"F        R   R   AAAAAAAAAAA   GGGG   IIIII   LLLLL   EEEEE"

		renderMenu := func(options []string, selected int) {
			for i, opt := range options {
				if i == selected {
					colors.GetColorPrinter("bartender").Print("  > " + opt + "\n")
				} else {
					colors.GetColorPrinter("bartender").Print("    " + opt + "\n")
				}
			}
		}

		err := keyboard.Open()
		if err != nil {
			TypeWriter("Error initializing keyboard input.", 20, colors.GetColorPrinter("system"))
			return ""
		}
		defer keyboard.Close()

		for {
			fmt.Print("\033[H")
			colors.GetColorPrinter("title").Print(asciiArt + "\n\n")
			renderMenu(menuOptions, selected)

			_, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			if key == keyboard.KeyArrowUp && selected > 0 {
				selected--
			} else if key == keyboard.KeyArrowDown && selected < len(menuOptions)-1 {
				selected++
			} else if key == keyboard.KeyEnter {
				switch selected {
				case 0:
					return "start"
				case 1:
					return "load"
				case 2:
					return "options"
				}
			}
		}
}
