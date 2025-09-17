package dialogue

import (
	"fmt"

	"github.com/michaelmastrucci/text-rpg/colors"
	"github.com/michaelmastrucci/text-rpg/utils"
)

// Character represents a game character with their dialogue
type Character struct {
	Name      string
	ColorFunc func(string) string
}

// Dialogue holds a piece of dialogue with its speaker and text
type Dialogue struct {
	Speaker *Character
	Text    string
}

// Scene represents a collection of dialogue lines for a specific scene
type Scene struct {
	ID        string
	Dialogues []Dialogue
}

// GameDialogues holds all dialogue in the game
type GameDialogues struct {
	Characters map[string]*Character
	Scenes     map[string]*Scene
}

// NewGameDialogues initializes a new GameDialogues instance
func NewGameDialogues() *GameDialogues {
	return &GameDialogues{
		Characters: make(map[string]*Character),
		Scenes:     make(map[string]*Scene),
	}
}

// AddCharacter adds a new character to the game
func (gd *GameDialogues) AddCharacter(id, name string, colorFunc func(string) string) {
	gd.Characters[id] = &Character{
		Name:      name,
		ColorFunc: colorFunc,
	}
}

// AddScene adds a new scene with its dialogues
func (gd *GameDialogues) AddScene(id string, dialogues []Dialogue) {
	gd.Scenes[id] = &Scene{
		ID:        id,
		Dialogues: dialogues,
	}
}

// GetScene returns a scene by ID
func (gd *GameDialogues) GetScene(id string) (*Scene, bool) {
	scene, exists := gd.Scenes[id]
	return scene, exists
}

// Format formats a dialogue line with the speaker's name and color
func (d *Dialogue) Format() string {
	if d.Speaker == nil {
		// For narrator text, just return the text with a border
		return utils.WrapTextInBorder(d.Text, 60)
	}
	
	// For character dialogue, include the speaker's name and apply their color
	speakerText := fmt.Sprintf("%s: %s", d.Speaker.Name, d.Text)
	if d.Speaker.ColorFunc != nil {
		speakerText = d.Speaker.ColorFunc(speakerText)
	}
	return utils.WrapTextInBorder(speakerText, 60)
}

// InitializeDialogues sets up all game dialogues
func InitializeDialogues() *GameDialogues {
	dialogues := NewGameDialogues()

	// Define characters
	dialogues.AddCharacter("bartender", "Bartender", func(s string) string {
		return colors.GetColorPrinter("bartender").Sprint(s)
	})
	dialogues.AddCharacter("narrator", "Narrator", func(s string) string {
		return colors.GetColorPrinter("narrator").Sprint(s)
	})
	dialogues.AddCharacter("player", "Player", func(s string) string {
		return colors.GetColorPrinter("player").Sprint(s)
	})

	// Define scenes
	dialogues.AddScene("intro", []Dialogue{
		{Speaker: dialogues.Characters["bartender"], Text: "Hey there and Welcome to-"},
		{Speaker: dialogues.Characters["bartender"], Text: "Oh sorry, I thought ya were a customer...yer the newbie, eh?"},
	})

	dialogues.AddScene("name_prompt", []Dialogue{
		{Speaker: dialogues.Characters["bartender"], Text: "Erm...Whats yer name again?"},
	})

	return dialogues
}
