package ui

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
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

		options := &colorm.DrawImageOptions{}
		options.GeoM.Scale(iconScale, iconScale)
		options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
		colorM := colorm.ColorM{}
		colorm.DrawImage(screen, roleImage, colorM, options)

	}
}

func drawTankSquads(screen *ebiten.Image, vd *ViewDimension, roleImages map[string]*ebiten.Image, sv *hll.ServerView, selectedPlayerID string) {
	for _, squad := range sv.Allies.Squads {
		if squad.SquadType == hll.StArmor {
			isSelected := squad.HasPlayer(selectedPlayerID)
			drawTankSquad(screen, vd, roleImages, squad, isSelected)
		}
	}
	for _, squad := range sv.Axis.Squads {
		if squad.SquadType == hll.StArmor {
			isSelected := squad.HasPlayer(selectedPlayerID)
			drawTankSquad(screen, vd, roleImages, squad, isSelected)
		}
	}
}

func drawTankSquad(screen *ebiten.Image, vd *ViewDimension, roleImages map[string]*ebiten.Image, squad *hll.SquadView, isSelected bool) {
	if squad == nil {
		return
	}

	if len(squad.Players) < 2 {
		return
	}

	drawTank := false
	tankPosition := hll.Position{}
	for i := 0; i < len(squad.Players); i++ {
		for j := i + 1; j < len(squad.Players); j++ {
			if squad.Players[i].IsSpawned() && squad.Players[i].Position == squad.Players[j].Position {
				tankPosition = squad.Players[i].Position
				drawTank = true
			}
		}
	}

	clr := CLR_ALLIES_LIGHT
	if squad.Team == hll.TmAxis {
		clr = CLR_AXIS_LIGHT
	}
	if isSelected {
		clr = CLR_SELECTED_LIGHT
	}

	if drawTank {
		x, y := util.TranslateCoords(vd.sizeX, vd.sizeY, tankPosition)
		x = x*vd.zoomLevel + vd.panX
		y = y*vd.zoomLevel + vd.panY

		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(vd.zoomLevel, TANK_SIZE_MODIFIER)), clr, false)

		roleImage, ok := roleImages[strings.ToLower(string(hll.TankCommander))]
		if ok {
			targetSize := util.IconSize(vd.zoomLevel, TANK_ICON_SIZE_MODIFIER)
			iconScale := targetSize / float64(roleImage.Bounds().Dx())

			options := &colorm.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			colorM := colorm.ColorM{}
			colorM.Scale(0, 0, 0, 1)
			colorm.DrawImage(screen, roleImage, colorM, options)
		}
	}

}
