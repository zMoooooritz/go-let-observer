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

func DrawGrid(screen *ebiten.Image, vd *shared.ViewDimension, currentMapID string, gameScore hll.TeamData) {
	width, height := vd.FrustumSize()

	cellWidth := width / float64(GRID_SIZE)
	cellHeight := height / float64(GRID_SIZE)

	// activeSectors := []int{1, 2, 1, 0, 0}

	currentMap := hll.MapToGameMap(hll.Map(currentMapID))

	for i := 0; i < GRID_SIZE; i++ {
		for j := 0; j < GRID_SIZE; j++ {
			x := float64(i)*cellWidth + vd.PanX
			y := float64(j)*cellHeight + vd.PanY

			if currentMap.Orientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				if currentMap.MirroredFactions {
					// Mirrored: Allies from right, Axis from left
					if GRID_SIZE-i <= gameScore.Allies {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
					}
					if i < gameScore.Axis {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
					}
				} else {
					// Default: Allies from left, Axis from right
					if i < gameScore.Allies {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
					}
					if GRID_SIZE-i <= gameScore.Axis {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
					}
				}

				// if activeSectors[i]+1 == j {
				// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ACTIVE_SECTOR_OVERLAY, false)
				// }
			}

			if currentMap.Orientation == hll.OriVertical {
				if i == 0 || i == 4 {
					continue
				}

				if currentMap.MirroredFactions {
					// Mirrored: Allies from bottom, Axis from top
					if GRID_SIZE-j <= gameScore.Allies {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
					}
					if j < gameScore.Axis {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
					}
				} else {
					// Default: Allies from top, Axis from bottom
					if j < gameScore.Allies {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_ALLIES_OVERLAY, false)
					}
					if GRID_SIZE-j <= gameScore.Axis {
						vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), shared.CLR_AXIS_OVERLAY, false)
					}
				}

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
