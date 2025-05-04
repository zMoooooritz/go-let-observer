package components

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func DrawPlayers(screen *ebiten.Image, vd *shared.ViewDimension, roleImages map[string]*ebiten.Image, playerList []hll.DetailedPlayerInfo, selectedPlayerID string) {
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

func drawPlayer(screen *ebiten.Image, vd *shared.ViewDimension, roleImages map[string]*ebiten.Image, player *hll.DetailedPlayerInfo, isSelected bool) {
	x, y := util.TranslateCoords(vd.SizeX, vd.SizeY, player.Position)
	x = x*vd.ZoomLevel + vd.PanX
	y = y*vd.ZoomLevel + vd.PanY

	sizeModifier := shared.PLAYER_SIZE_MODIFIER
	clr := shared.CLR_ALLIES
	if player.Team == hll.TmAxis {
		clr = shared.CLR_AXIS
	}
	if isSelected {
		sizeModifier = shared.SELECTED_PLAYER_SIZE_MODIFIER
		clr = shared.CLR_SELECTED
	}

	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(util.IconCircleRadius(vd.ZoomLevel, sizeModifier)), clr, false)

	roleImage, ok := roleImages[strings.ToLower(string(player.Role))]
	if ok {
		targetSize := util.IconSize(vd.ZoomLevel, sizeModifier)
		iconScale := targetSize / float64(roleImage.Bounds().Dx())

		options := &colorm.DrawImageOptions{}
		options.GeoM.Scale(iconScale, iconScale)
		options.GeoM.Translate(x-targetSize/2, y-targetSize/2)
		colorM := colorm.ColorM{}
		colorm.DrawImage(screen, roleImage, colorM, options)

	}
}
