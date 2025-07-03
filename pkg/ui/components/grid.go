package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	MAIN_GRID_STROKE_WIDTH = 2
	SUB_GRID_STROKE_WIDTH  = 1

	GRID_SIZE = 5
)

var (
	GRID_COLOR = color.RGBA{100, 100, 100, 255}
	FILL_COLOR = color.RGBA{50, 50, 50, 100}
)

func DrawGrid(screen *ebiten.Image, vd *shared.ViewDimension, orientation hll.Orientation) {
	width, height := vd.FrustumSize()

	cellWidth := width / float64(GRID_SIZE)
	cellHeight := height / float64(GRID_SIZE)

	// gameScore := hll.TeamData{
	// 	Allies: 2,
	// 	Axis:   3,
	// }

	// activeSectors := []int{1, 2, 1, 0, 0}

	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			x := float64(i)*cellWidth + vd.PanX
			y := float64(j)*cellHeight + vd.PanY

			if orientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				// if i < gameScore.Allies {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
				// }
				// if gridSize-i <= gameScore.Axis {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
				// }

				// if activeSectors[i]+1 == j {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ACTIVE_SECTOR_OVERLAY, false)
				// }
			}

			if orientation == hll.OriVertical {
				if i == 0 || i == 4 {
					continue
				}

				// if j < gameScore.Allies {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
				// }
				// if gridSize-j <= gameScore.Axis {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
				// }

				// if activeSectors[j]+1 == i {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ACTIVE_SECTOR_OVERLAY, false)
				// }
			}

			strokeWidth := util.AdaptiveScaledDim(MAIN_GRID_STROKE_WIDTH, vd.ZoomLevel)

			// main grid lines
			vector.StrokeLine(screen, float32(x), float32(y), float32(x+cellWidth), float32(y), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x), float32(y+cellHeight), float32(x+cellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x), float32(y), float32(x), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x+cellWidth), float32(y), float32(x+cellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)

			subCellWidth := cellWidth / 2
			subCellHeight := cellHeight / 2

			strokeWidth = util.AdaptiveScaledDim(SUB_GRID_STROKE_WIDTH, vd.ZoomLevel)

			// sub grid lines
			vector.StrokeLine(screen, float32(x), float32(y+subCellHeight), float32(x+cellWidth), float32(y+subCellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x+subCellWidth), float32(y), float32(x+subCellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
		}
	}
}
