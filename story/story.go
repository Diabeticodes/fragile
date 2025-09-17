package story

import (
	"fmt"

	"github.com/michaelmastrucci/text-rpg/colors"
	"github.com/michaelmastrucci/text-rpg/ui"
	"github.com/michaelmastrucci/text-rpg/utils"
)

// ShowIntroStory displays the game's introduction story
func bartenderDialog() {
	// Clear the screen before showing the story
	ui.ClearScreen()

	// Story paragraphs with proper spacing
	paragraphs := []string{
		"A once bustling city, filled with life and laughter, has become a hollow shell of its former self.",
		"It's cobblestone streets used to echo with conversations...",
		"And even it's darkest nights were illuminated by streetlamps and lively establishments that reflected off the eyes of it's inhabitants-",
		"creating an enchanting splendor that guided you on your way back to wherever you called home.",
		"Then one night, something dreadful happened to the peaceful city...",
		"The joyful noise faded into silence, and the alleys became quiet",
		"The alleys are empty, the silence uncomfortable, and the avenues lined with fog..."
		

	}

	// Display each paragraph with a typewriter effect and delay between them
	for _, p := range paragraphs {
		ui.TypeWriter("\n"+p+"\n", 20, colors.GetColorPrinter("bartender"))
		utils.Delay(3) // Pause between paragraphs
	}

	// Wait for user to press enter to continue
	ui.TypeWriter("\nPress [Enter] to begin your first day...", 20, colors.GetColorPrinter("input"))
	var input string
	fmt.Scanln(&input)
}
