package components

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
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
		{"T", "Toggle tanks/spas"},
		{"G", "Toggle grid overlay"},
		{"H", "Toggle header overlay"},
		{"Tab", "Show scoreboard"},
		{"Space", "Toggle replay pause"},
		{"Right", "Seek forward in replay"},
		{"Left", "Seek backward in replay"},
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

	helpCache *ebiten.Image
)

const (
	HELP_WIDTH  = 600
	HELP_HEIGHT = 400
)

func DrawHelp(screen *ebiten.Image) {
	if helpCache == nil {
		helpCache = util.NewScaledImage(HELP_WIDTH, HELP_HEIGHT)

		util.DrawScaledRect(helpCache, 0, 0, HELP_WIDTH, HELP_HEIGHT, shared.CLR_OVERLAY)

		textX := 20
		textY := 30
		lineHeight := 30
		util.DrawText(helpCache, "Help", textX, textY, shared.CLR_WHITE, util.Font.Normal)
		textY += lineHeight
		textX += 20

		for _, mouseaction := range mouseactions {
			util.DrawText(helpCache, formatHelpLine(mouseaction.Action, mouseaction.Key), textX, textY, shared.CLR_WHITE, util.Font.Small)
			textY += 20
		}

		textY += 20

		for _, keybind := range keybinds {
			util.DrawText(helpCache, formatHelpLine(keybind.Action, keybind.Key), textX, textY, shared.CLR_WHITE, util.Font.Small)
			textY += 20
		}
	}

	screenWidth := shared.ROOT_SCALING_SIZE
	screenHeight := shared.ROOT_SCALING_SIZE
	helpX := (screenWidth - HELP_WIDTH) / 2
	helpY := (screenHeight - HELP_HEIGHT) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(util.ScaledDim(helpX)), float64(util.ScaledDim(helpY)))
	screen.DrawImage(helpCache, op)
}

func formatHelpLine(action, key string) string {
	return fmt.Sprintf("%-25s : %s", key, action)
}
