package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	keybinds = []struct {
		Key    string
		Action string
	}{
		{"+ / -", "Increase/Decrease update interval"},
		{"P", "Toggle players"},
		{"I", "Toggle player info"},
		{"S", "Toggle guesstimated spawns"},
		{"G", "Toggle grid overlay"},
		{"H", "Toggle header overlay"},
		{"Tab", "Show scoreboard"},
		{"?", "Show this help"},
		{"Esc, Q, Ctrl+C", "Exit the application"},
	}

	mouseactions = []struct {
		Key    string
		Action string
	}{
		{"LeftClick", "Select player"},
		{"RightClick-Drag", "Pan the map"},
		{"MouseWheel", "Zoom in/out on the map"},
	}
)

func (mv *MapView) drawHelp(screen *ebiten.Image) {
	helpWidth := 600
	helpHeight := 400
	screenWidth := ROOT_SCALING_SIZE
	screenHeight := ROOT_SCALING_SIZE
	helpX := (screenWidth - helpWidth) / 2
	helpY := (screenHeight - helpHeight) / 2

	util.DrawScaledRect(screen, helpX, helpY, helpWidth, helpHeight, CLR_OVERLAY)

	textX := helpX + 20
	textY := helpY + 40
	lineHeight := 30
	util.DrawText(screen, "Help", textX, textY, CLR_WHITE, util.Font.Normal)
	textY += lineHeight
	textX += 20

	for _, mouseaction := range mouseactions {
		util.DrawText(screen, formatHelpLine(mouseaction.Action, mouseaction.Key), textX, textY, CLR_WHITE, util.Font.Small)
		textY += 20
	}

	textY += 20

	for _, keybind := range keybinds {
		util.DrawText(screen, formatHelpLine(keybind.Action, keybind.Key), textX, textY, CLR_WHITE, util.Font.Small)
		textY += 20
	}
}

func formatHelpLine(action, key string) string {
	return fmt.Sprintf("%-25s : %s", key, action)
}
