package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type ViewDimension struct {
	sizeX     int
	sizeY     int
	zoomLevel float64
	panX      float64
	panY      float64
}

func (vd *ViewDimension) FrustumSize() (float64, float64) {
	screenSizeX := float64(vd.sizeX) * vd.zoomLevel
	screenSizeY := float64(vd.sizeY) * vd.zoomLevel
	return screenSizeX, screenSizeY
}

type BaseViewer struct {
	backgroundImage *ebiten.Image
	dim             *ViewDimension
	ctx             StateContext
}

func NewBaseViewer(ctx StateContext) *BaseViewer {
	return &BaseViewer{
		backgroundImage: nil,
		dim: &ViewDimension{
			sizeX:     util.Config.UIOptions.ScreenSize,
			sizeY:     util.Config.UIOptions.ScreenSize,
			zoomLevel: MIN_ZOOM_LEVEL,
			panX:      0.0,
			panY:      0.0,
		},
		ctx: ctx,
	}
}

func (bv *BaseViewer) DrawBackground(screen *ebiten.Image) {
	if bv.backgroundImage != nil {
		screenSize := screen.Bounds().Size()
		imageSize := bv.backgroundImage.Bounds().Size()
		imageScale := float64(screenSize.X) / float64(imageSize.X)

		if bv.dim.zoomLevel != 0 {
			imageScale *= bv.dim.zoomLevel
		}

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(imageScale, imageScale)
		if bv.dim.panX != 0 || bv.dim.panY != 0 {
			options.GeoM.Translate(bv.dim.panX, bv.dim.panY)
		}
		screen.DrawImage(bv.backgroundImage, options)
	} else {
		screen.Fill(FALLBACK_BACKGROUND)
	}
}

func (bv *BaseViewer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return bv.dim.sizeX, bv.dim.sizeY
}
