package views

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type BaseViewer struct {
	backgroundImage *ebiten.Image
	dim             *shared.ViewDimension
	ctx             shared.StateContext
}

func NewBaseViewer(ctx shared.StateContext) *BaseViewer {
	return &BaseViewer{
		backgroundImage: nil,
		dim: &shared.ViewDimension{
			SizeX:     util.Config.UIOptions.ScreenSize,
			SizeY:     util.Config.UIOptions.ScreenSize,
			ZoomLevel: shared.MIN_ZOOM_LEVEL,
			PanX:      0.0,
			PanY:      0.0,
		},
		ctx: ctx,
	}
}

func (bv *BaseViewer) DrawBackground(screen *ebiten.Image) {
	if bv.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := bv.backgroundImage.Bounds().Size()
		imageScale := float64(screenSize.X) / float64(imageSize.X)

		if bv.dim.ZoomLevel != 0 {
			imageScale *= bv.dim.ZoomLevel
		}

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(imageScale, imageScale)
		if bv.dim.PanX != 0 || bv.dim.PanY != 0 {
			options.GeoM.Translate(bv.dim.PanX, bv.dim.PanY)
		}
		screen.DrawImage(bv.backgroundImage, options)
	} else {
		screen.Fill(shared.FALLBACK_BACKGROUND)
	}
}

func (bv *BaseViewer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return bv.dim.SizeX, bv.dim.SizeY
}
