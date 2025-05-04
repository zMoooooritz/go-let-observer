package components

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

var (
	cachedServerName      *ebiten.Image
	lastServerName        time.Time
	cachedPlayerCount     *ebiten.Image
	lastPlayerCountUpdate time.Time
)

func DrawServerName(screen *ebiten.Image, serverName string) {
	currentTime := time.Now()
	if cachedServerName == nil || currentTime.Sub(lastServerName) >= time.Second {
		overlayWidth := 750
		overlayHeight := 50
		cachedServerName = ebiten.NewImage(overlayWidth, overlayHeight)

		util.DrawScaledRect(cachedServerName, 0, 0, overlayWidth, overlayHeight, shared.CLR_OVERLAY)

		textX := 10
		textY := 30
		util.DrawText(cachedServerName, serverName, textX, textY, shared.CLR_WHITE, util.Font.Normal)
		lastServerName = currentTime
	}

	screen.DrawImage(cachedServerName, nil)
}

func DrawPlayerCount(screen *ebiten.Image, playerCurrCount, playerMaxCount int) {
	currentTime := time.Now()
	if cachedPlayerCount == nil || currentTime.Sub(lastPlayerCountUpdate) >= time.Second {
		overlayWidth := 200
		overlayHeight := 50
		cachedPlayerCount = ebiten.NewImage(overlayWidth, overlayHeight)

		util.DrawScaledRect(cachedPlayerCount, 0, 0, overlayWidth, overlayHeight, shared.CLR_OVERLAY)

		textX := 10
		textY := 30
		info := fmt.Sprintf("Players: %d/%d", playerCurrCount, playerMaxCount)
		util.DrawText(cachedPlayerCount, info, textX, textY, shared.CLR_WHITE, util.Font.Normal)
		lastPlayerCountUpdate = currentTime
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 50)
	screen.DrawImage(cachedPlayerCount, op)
}
