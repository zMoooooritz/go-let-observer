package components

import (
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func DrawTankSquads(screen *ebiten.Image, vd *shared.ViewDimension, roleImages map[string]*ebiten.Image, sv *hll.ServerView, selectedPlayerID string) {
	// Sort allies squads by name
	alliesSquadNames := make([]string, 0, len(sv.Allies.Squads))
	for name := range sv.Allies.Squads {
		alliesSquadNames = append(alliesSquadNames, name)
	}
	slices.Sort(alliesSquadNames)

	for _, name := range alliesSquadNames {
		squad := sv.Allies.Squads[name]
		if squad.SquadType == hll.StArmor {
			isSelected := squad.HasPlayer(selectedPlayerID)
			drawTankSquad(screen, vd, roleImages, squad, isSelected)
		}
	}

	// Sort axis squads by name
	axisSquadNames := make([]string, 0, len(sv.Axis.Squads))
	for name := range sv.Axis.Squads {
		axisSquadNames = append(axisSquadNames, name)
	}
	slices.Sort(axisSquadNames)

	for _, name := range axisSquadNames {
		squad := sv.Axis.Squads[name]
		if squad.SquadType == hll.StArmor {
			isSelected := squad.HasPlayer(selectedPlayerID)
			drawTankSquad(screen, vd, roleImages, squad, isSelected)
		}
	}
}

func drawTankSquad(screen *ebiten.Image, vd *shared.ViewDimension, roleImages map[string]*ebiten.Image, squad *hll.SquadView, isSelected bool) {
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

	clr := shared.CLR_ALLIES_LIGHT
	if squad.Team == hll.TmAxis {
		clr = shared.CLR_AXIS_LIGHT
	}
	if isSelected {
		clr = shared.CLR_SELECTED_LIGHT
	}

	if drawTank {
		x, y := util.TranslateCoords(vd.SizeX, vd.SizeY, tankPosition)
		x = x*vd.ZoomLevel + vd.PanX
		y = y*vd.ZoomLevel + vd.PanY

		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(vd.ZoomLevel, shared.TANK_SIZE_MODIFIER)), clr, false)

		roleImage, ok := roleImages[strings.ToLower(string(hll.TankCommander))]
		if ok {
			targetSize := util.IconSize(vd.ZoomLevel, shared.TANK_ICON_SIZE_MODIFIER)
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
