package util

import (
	"log"

	"github.com/zMoooooritz/go-let-observer/assets"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	Huge   font.Face
	Title  font.Face
	Normal font.Face
	Small  font.Face
}

func scaledFontSize(fontSize int, screenSize int) float64 {
	return float64(fontSize) * float64(screenSize) / 1000
}

func LoadFonts(screenSize int) Font {
	fnt := Font{}

	fontData, err := assets.Assets.ReadFile("fonts/RobotoMono-Regular.ttf")
	if err != nil {
		log.Fatalf("failed to load font: %v", err)
	}

	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	smallFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    scaledFontSize(10, screenSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create normal font: %v", err)
	}
	fnt.Small = smallFont

	normalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    scaledFontSize(14, screenSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create normal font: %v", err)
	}
	fnt.Normal = normalFont

	titleFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    scaledFontSize(18, screenSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create.Title font: %v", err)
	}
	fnt.Title = titleFont

	hugeFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    scaledFontSize(24, screenSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create.Title font: %v", err)
	}
	fnt.Huge = hugeFont

	return fnt
}
