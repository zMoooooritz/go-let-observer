package main

import (
	"log"

	"github.com/zMoooooritz/go-let-observer/assets"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	huge   font.Face
	title  font.Face
	normal font.Face
	small  font.Face
}

func loadFonts() Font {
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
		Size:    10,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create normal font: %v", err)
	}
	fnt.small = smallFont

	normalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create normal font: %v", err)
	}
	fnt.normal = normalFont

	titleFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create title font: %v", err)
	}
	fnt.title = titleFont

	hugeFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create title font: %v", err)
	}
	fnt.huge = hugeFont

	return fnt
}
