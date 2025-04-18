package ui

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func drawPlayers(screen *ebiten.Image, vd *ViewDimension, roleImages map[string]*ebiten.Image, playerList []hll.DetailedPlayerInfo, selectedPlayerID string) {
	var selectedPlayer *hll.DetailedPlayerInfo

	for _, player := range playerList {
		if !player.IsSpawned() {
			continue
		}

		if selectedPlayerID != "" && player.ID == selectedPlayerID {
			selectedPlayer = &player
			continue
		}

		drawPlayer(screen, vd, roleImages, &player, false)
	}

	if selectedPlayer != nil {
		drawPlayer(screen, vd, roleImages, selectedPlayer, true)
	}
}

func drawPlayer(screen *ebiten.Image, vd *ViewDimension, roleImages map[string]*ebiten.Image, player *hll.DetailedPlayerInfo, isSelected bool) {
	x, y := util.TranslateCoords(vd.sizeX, vd.sizeY, player.Position)
	x = x*vd.zoomLevel + vd.panX
	y = y*vd.zoomLevel + vd.panY

	sizeModifier := PLAYER_SIZE_MODIFIER
	clr := CLR_ALLIES
	if player.Team == hll.TmAxis {
		clr = CLR_AXIS
	}
	if isSelected {
		sizeModifier = SELECTED_PLAYER_SIZE_MODIFIER
		clr = CLR_SELECTED
	}

	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(vd.zoomLevel, sizeModifier)), clr, false)

	roleImage, ok := roleImages[strings.ToLower(string(player.Role))]
	if ok {
		targetSize := util.IconSize(vd.zoomLevel, sizeModifier)
		iconScale := targetSize / float64(roleImage.Bounds().Dx())

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(iconScale, iconScale)
		options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
		screen.DrawImage(roleImage, options)
	}
}
