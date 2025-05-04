package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func DrawSpawns(screen *ebiten.Image, spawns []rcndata.SpawnPoint, spawnImages map[string]*ebiten.Image, vd *shared.ViewDimension) {
	for _, spawn := range spawns {
		if spawn.SpawnType == rcndata.SpawnTypeNone {
			continue
		}

		x, y := util.TranslateCoords(vd.SizeX, vd.SizeY, spawn.Position)
		x = x*vd.ZoomLevel + vd.PanX
		y = y*vd.ZoomLevel + vd.PanY

		clr := shared.CLR_ALLIES_DARK
		if spawn.Team == hll.TmAxis {
			clr = shared.CLR_AXIS_DARK
		}

		rectSize := int(2 * util.IconCircleRadius(vd.ZoomLevel, shared.SPAWN_SIZE_MODIFIER))
		util.DrawScaledRect(screen, int(x)-rectSize/2, int(y)-rectSize/2, rectSize, rectSize, clr)

		spawnImage, ok := spawnImages[string(spawn.SpawnType)]
		if ok {
			targetSize := util.IconSize(vd.ZoomLevel, shared.SPAWN_SIZE_MODIFIER)
			iconScale := targetSize / float64(spawnImage.Bounds().Dx())

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			screen.DrawImage(spawnImage, options)
		}
	}
}
