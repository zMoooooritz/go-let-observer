package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func drawSpawns(screen *ebiten.Image, spawns []rcndata.SpawnPoint, spawnImages map[string]*ebiten.Image, vd *ViewDimension) {
	for _, spawn := range spawns {
		if spawn.SpawnType == rcndata.SpawnTypeNone {
			continue
		}

		x, y := util.TranslateCoords(vd.sizeX, vd.sizeY, spawn.Position)
		x = x*vd.zoomLevel + vd.panX
		y = y*vd.zoomLevel + vd.panY

		clr := CLR_ALLIES_SPAWN
		if spawn.Team == hll.TmAxis {
			clr = CLR_AXIS_SPAWN
		}

		rectSize := int(2 * util.IconCircleRadius(vd.zoomLevel, SPAWN_SIZE_MODIFIER))
		util.DrawScaledRect(screen, int(x)-rectSize/2, int(y)-rectSize/2, rectSize, rectSize, clr)

		spawnImage, ok := spawnImages[string(spawn.SpawnType)]
		if ok {
			targetSize := util.IconSize(vd.zoomLevel, SPAWN_SIZE_MODIFIER)
			iconScale := targetSize / float64(spawnImage.Bounds().Dx())

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(iconScale, iconScale)
			options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
			screen.DrawImage(spawnImage, options)
		}
	}
}
