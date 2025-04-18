package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	MAIN_GRID_STROKE_WIDTH = 2
	SUB_GRID_STROKE_WIDTH  = 1
)

var (
	GRID_COLOR = color.RGBA{100, 100, 100, 255}
	FILL_COLOR = color.RGBA{50, 50, 50, 100}
)

func drawGrid(screen *ebiten.Image, vd *ViewDimension, orientation hll.Orientation) {
	width, height := vd.FrustumSize()

	cellWidth := width / 5
	cellHeight := height / 5

	active := []int{1, 2, 1, 0, 0}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := float64(i)*cellWidth + vd.panX
			y := float64(j)*cellHeight + vd.panY

			if orientation == hll.OriHorizontal {
				if j == 0 || j == 4 {
					continue
				}

				if active[i]+1 == j {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			if orientation == hll.OriVertical {
				if i == 0 || i == 4 {
					continue
				}

				if active[j]+1 == i {
					// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), fillColor, false)
				}
			}

			strokeWidth := util.AdaptiveScaledDim(MAIN_GRID_STROKE_WIDTH, vd.zoomLevel)

			// main grid lines
			vector.StrokeLine(screen, float32(x), float32(y), float32(x+cellWidth), float32(y), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x), float32(y+cellHeight), float32(x+cellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x), float32(y), float32(x), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x+cellWidth), float32(y), float32(x+cellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)

			subCellWidth := cellWidth / 2
			subCellHeight := cellHeight / 2

			strokeWidth = util.AdaptiveScaledDim(SUB_GRID_STROKE_WIDTH, vd.zoomLevel)

			// sub grid lines
			vector.StrokeLine(screen, float32(x), float32(y+subCellHeight), float32(x+cellWidth), float32(y+subCellHeight), float32(strokeWidth), GRID_COLOR, false)
			vector.StrokeLine(screen, float32(x+subCellWidth), float32(y), float32(x+subCellWidth), float32(y+cellHeight), float32(strokeWidth), GRID_COLOR, false)
		}
	}
}
