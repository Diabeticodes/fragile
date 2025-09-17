package main

import (
	"fmt"
	"strings"

	"github.com/text-rpg/colors"
	"github.com/text-rpg/dialogue"
	"github.com/text-rpg/ui"
	"github.com/text-rpg/utils"
)

type Player struct {
	Name      string
	Level     int
	XP        int
	Recipes   []string
	Inventory []string
}

var gameDialogues *dialogue.GameDialogues


func init() {
	// Attempts to set terminal to 30 rows, 98 columns
	fmt.Print("\033[8;30;98t") 

	// Initialize all game dialogues
	gameDialogues = dialogue.InitializeDialogues()
}

func showGameIntro() {
	
	// Clear screen and show title
	ui.ClearScreen()

	// Get the intro scene
	introScene, exists := gameDialogues.GetScene("intro")
	if !exists {
		// Fallback in case the scene is not found
		ui.TypeWriter("Welcome to Fragile!", 50, colors.GetColorPrinter("narrator"))
		return
	}

	// Collect all dialogue text
	var allText strings.Builder
	for _, d := range introScene.Dialogues {
		// Add speaker name if available
		if d.Speaker != nil {
			allText.WriteString(d.Speaker.Name + ":\n")
		}
		allText.WriteString(d.Text + "\n\n")
	}

	// Wrap all text in a single border
	formattedText := utils.WrapTextInBorder(allText.String(), 70)
	
	// Print the formatted text with border
	fmt.Println(formattedText)
}

func startNewGame() {
		ui.TopLocationBar("Edge of the Table - Bar", "00:00")
	// Get the name prompt scene
	nameScene, exists := gameDialogues.GetScene("name_prompt")
	if exists && len(nameScene.Dialogues) > 0 {
		// Show the first line of the name prompt
		d := nameScene.Dialogues[0]
		ui.TypeWriter(d.Text, 50, colors.GetColorPrinter("bartender"))
		utils.Delay(1)
	}

	// Get player name with colored prompt
	colors.GetColorPrinter("input").Print("Please enter your name: ")
	var name string
	fmt.Scan(&name)

	player := Player{
		Name:     name,
		Level:    1,
		XP:       0,
		Recipes:  []string{"Beer", "Wine", "Whiskey", "Vodka"},
		Inventory: []string{},
	}
    
    // Continue with the rest of your game...
    // Get the new game scene
    newGameScene, exists := gameDialogues.GetScene("new_game")
    if exists {
        // Display each line of dialogue in the scene
        for _, d := range newGameScene.Dialogues {
            ui.TypeWriter(d.Text, 50, colors.GetColorPrinter("bartender"))
            utils.Delay(1)
        }
    } else {
        // Fallback in case the scene is not found
        ui.TypeWriter(fmt.Sprintf("\nSo yer, %s? Nice ta meet ya.", player.Name), 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("Before ya jump behind the counter 'n start throwin' my booze everywhere, lemme lay some ground rules.", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("Since yer new, I gotta start ya at level "+fmt.Sprint(player.Level)+" since ya got "+fmt.Sprint(player.XP)+" XP.", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("It's real simple, see. Customer's come in fer drinks. It's YER job to make 'em and make 'em RIGHT.", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("Ya can make "+fmt.Sprint(len(player.Recipes))+" different drinks...", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter(strings.Join(player.Recipes, ", "), 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("Well i's a bit less ammo than I hoped you'd be packin'...but I got a feelin' you'll pick things up pretty quick.", 50, colors.GetColorPrinter("bartender"))

        ui.TypeWriter("The recipe book is under the counter. Open it any time by pressing 'R' to make sure ya don't mess up an order, Yeah?", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("The ingredients are behind you on the wall. You can check supplies by pressing 'I' to make sure YA DON'T MESS UP AN ORDER. YEAH?", 50, colors.GetColorPrinter("bartender"))
        ui.TypeWriter("The bar is in front of you. You can check the bar by pressing 'B'", 50, colors.GetColorPrinter("bartender"))
    }
}

func showOptions() {
    ui.ClearScreen()
    ui.TypeWriter("=== GAME OPTIONS ===\n", 20, colors.GetColorPrinter("title"))
    ui.TypeWriter("1. Sound: ON/OFF", 20, colors.GetColorPrinter("bartender"))
    ui.TypeWriter("2. Text Speed: NORMAL", 20, colors.GetColorPrinter("bartender"))
    ui.TypeWriter("3. Back to Main Menu [B]\n", 20, colors.GetColorPrinter("bartender"))
    
    // Simple implementation - just wait for any key to return
    colors.GetColorPrinter("input").Print("Press any key to return to the main menu...")
    var input string
    fmt.Scanln(&input)
}

func main() {
    // Show start screen and get user's choice
    for {
        choice := ui.ShowStartScreen()
        
        switch choice {
        case "start":
            showGameIntro()
            startNewGame()
            return
        case "load":
            ui.ClearScreen()
            ui.TypeWriter("Load game functionality coming soon!", 20, colors.GetColorPrinter("system"))
            utils.Delay(2)
        case "options":
            showOptions()
        }
    }
}