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

var vehicleSquadImages = map[hll.SquadType]hll.Role{
	hll.StArmor:     hll.TankCommander,
	hll.StArtillery: hll.ArtilleryObserver,
}

func DrawVehicleSquads(screen *ebiten.Image, vd *shared.ViewDimension, roleImages map[string]*ebiten.Image, sv *hll.ServerView, selectedPlayerID string) {
	// Sort allies squads by name
	alliesSquadNames := make([]string, 0, len(sv.Allies.Squads))
	for name := range sv.Allies.Squads {
		alliesSquadNames = append(alliesSquadNames, name)
	}
	slices.Sort(alliesSquadNames)

	for _, name := range alliesSquadNames {
		squad := sv.Allies.Squads[name]
		if squadImageRole, ok := vehicleSquadImages[squad.SquadType]; ok {
			squadImage, ok := roleImages[strings.ToLower(string(squadImageRole))]
			if ok {
				isSelected := squad.HasPlayer(selectedPlayerID)
				drawVehicleSquad(screen, vd, squadImage, squad, isSelected)
			}
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
		if squadImageRole, ok := vehicleSquadImages[squad.SquadType]; ok {
			squadImage, ok := roleImages[strings.ToLower(string(squadImageRole))]
			if ok {
				isSelected := squad.HasPlayer(selectedPlayerID)
				drawVehicleSquad(screen, vd, squadImage, squad, isSelected)
			}
		}
	}
}

func drawVehicleSquad(screen *ebiten.Image, vd *shared.ViewDimension, squadImage *ebiten.Image, squad *hll.SquadView, isSelected bool) {
	if squad == nil {
		return
	}

	if len(squad.Players) < 2 {
		return
	}

	drawVehicle := false
	vehiclePosition := hll.Position{}
	for i := 0; i < len(squad.Players); i++ {
		for j := i + 1; j < len(squad.Players); j++ {
			if squad.Players[i].IsSpawned() && squad.Players[i].Position == squad.Players[j].Position {
				vehiclePosition = squad.Players[i].Position
				drawVehicle = true
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

	if drawVehicle {
		x, y := util.TranslateCoords(vd.SizeX, vd.SizeY, vehiclePosition)
		x = x*vd.ZoomLevel + vd.PanX
		y = y*vd.ZoomLevel + vd.PanY

		vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(vd.ZoomLevel, shared.VEHICLE_SIZE_MODIFIER)), clr, false)

		targetSize := util.IconSize(vd.ZoomLevel, shared.VEHICLE_ICON_SIZE_MODIFIER)
		iconScale := targetSize / float64(squadImage.Bounds().Dx())

		options := &colorm.DrawImageOptions{}
		options.GeoM.Scale(iconScale, iconScale)
		options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
		colorM := colorm.ColorM{}
		colorM.Scale(0, 0, 0, 1)
		colorm.DrawImage(screen, squadImage, colorM, options)
	}

}
