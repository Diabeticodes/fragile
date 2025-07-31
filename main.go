package main

import (
	"fmt"
	"strings"

	"github.com/michaelmastrucci/text-rpg/colors"
	"github.com/michaelmastrucci/text-rpg/story"
	"github.com/michaelmastrucci/text-rpg/ui"
	"github.com/michaelmastrucci/text-rpg/utils"
)

type Player struct {
	Name     string
	Level    int
	XP       int
	Recipes  []string
	Inventory []string
}

func showGameIntro() {
	// Show the story first
	story.ShowIntroStory()
	
	// Then show the game title and introduction
	ui.ClearScreen()
	gameTitle := colors.Colorize("Edge of the Table", "title")
    fmt.Println(gameTitle)
    
    // Welcome messages with typewriter effect and colors
    ui.TypeWriter("\nHey there and Welcome to-", 50, colors.GetColorPrinter("bartender"))
    utils.Delay(1)
    ui.TypeWriter("Oh sorry, I thought ya were a customer...yer the newbie, eh?", 50, colors.GetColorPrinter("bartender"))
    utils.Delay(2)
}

func startNewGame() {
    ui.TypeWriter("Erm...Whats yer name again?", 50, colors.GetColorPrinter("bartender"))
    utils.Delay(1)
    
    // Get player name with colored prompt
    colors.GetColorPrinter("input").Print("Please enter your name: ")
    var name string
    fmt.Scan(&name)

    player := Player{
        Name: name,
        Level: 1,
        XP: 0,
        Recipes: []string{"Beer", "Wine", "Whiskey", "Vodka"},
    }
    
    // Continue with the rest of your game...
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