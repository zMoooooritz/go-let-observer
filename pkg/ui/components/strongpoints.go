package components

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	STRONG_POINT_STROKE_WIDTH = 3
	STRIPE_SPACING            = 5.0
	STRIPE_WIDTH              = 1.0
)

func drawDiagonalStripes(screen *ebiten.Image, centerX, centerY, radius float32, stripeSpacing, stripeWidth float32, color color.Color) {
	// Draw diagonal stripes from top-right to bottom-left at -45 degrees
	// For a -45-degree line: y - centerY = -(x - centerX), or y = -x + centerX + centerY
	// We'll iterate through lines parallel to this

	// The perpendicular distance from center to a parallel line offset by d is d/sqrt(2)
	maxOffset := radius * float32(math.Sqrt2)

	for offset := -maxOffset; offset <= maxOffset; offset += stripeSpacing {
		// For a line parallel to y = -x, offset by 'offset' in perpendicular direction

		// Calculate perpendicular distance from center to this line
		perpDist := float32(math.Abs(float64(offset)))

		// If line doesn't intersect circle, skip it
		if perpDist > radius {
			continue
		}

		// Calculate half-chord length
		halfChord := float32(math.Sqrt(float64(radius*radius - perpDist*perpDist)))

		// The line intersects the circle at two points
		// Moving along the line direction (1/sqrt(2), -1/sqrt(2)) - right and up
		dx := halfChord / float32(math.Sqrt2)
		dy := -halfChord / float32(math.Sqrt2)

		// Point on the line closest to center (perpendicular from center)
		// For -45-degree line, perpendicular direction is (1/sqrt(2), 1/sqrt(2))
		midX := centerX + offset/float32(math.Sqrt2)
		midY := centerY + offset/float32(math.Sqrt2)

		// Calculate intersection points
		x1 := midX - dx
		y1 := midY - dy
		x2 := midX + dx
		y2 := midY + dy

		// Draw the clipped line segment
		vector.StrokeLine(screen, x1, y1, x2, y2, stripeWidth, color, false)
	}
}

func DrawStrongpoints(screen *ebiten.Image, vd *shared.ViewDimension, currentLayerID hll.LayerIdentifier, selectedSectors [5]int) {
	layer := currentLayerID.Layer()

	for i, sectorIndex := range selectedSectors {
		if sectorIndex == 0 {
			continue
		}

		sectors, err := layer.Sectors()
		if err != nil {
			continue
		}
		captureZones := sectors[i].CaptureZones
		if sectorIndex > len(captureZones) {
			continue
		}

		sector := captureZones[sectorIndex-1]

		centerX, centerY := util.TranslateCoords(vd.SizeX, vd.SizeY, sector.Strongpoint.Center)
		centerX = centerX*vd.ZoomLevel + vd.PanX
		centerY = centerY*vd.ZoomLevel + vd.PanY

		radius := util.TranslateDistance(vd.SizeX, vd.SizeY, sector.Strongpoint.Radius)
		radius = radius * vd.ZoomLevel

		strokeWidth := util.AdaptiveScaledDim(STRONG_POINT_STROKE_WIDTH, vd.ZoomLevel)
		stripeWidth := util.AdaptiveScaledDim(STRIPE_WIDTH, vd.ZoomLevel)
		stripeSpacing := util.AdaptiveScaledDim(STRIPE_SPACING, vd.ZoomLevel)
		color := shared.CLR_BLACK

		// Draw diagonal stripes
		drawDiagonalStripes(screen, float32(centerX), float32(centerY), float32(radius), float32(stripeSpacing), float32(stripeWidth), color)

		// Draw circle outline
		vector.StrokeCircle(screen, float32(centerX), float32(centerY), float32(radius), float32(strokeWidth), color, false)
	}
}
