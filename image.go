package main

import (
	"bytes"
	"image"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
)

const (
	roleCount = 14
)

func loadRoleImages() map[string]*ebiten.Image {
	roleImages := make(map[string]*ebiten.Image)
	for index := range roleCount {
		roleName := strings.ToLower(string(hll.RoleFromInt(index)))
		imgData, err := assets.Assets.ReadFile("roles/" + roleName + ".png")
		if err != nil {
			log.Printf("Error loading role image for %s: %v", roleName, err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			log.Printf("Error loading role image for %s: %v", roleName, err)
			continue
		}
		roleImages[roleName] = ebiten.NewImageFromImage(img)
	}
	return roleImages
}

func loadMapImage(mapName string) (*ebiten.Image, error) {
	imgData, err := assets.Assets.ReadFile("tacmaps/" + mapName + ".png")
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func loadGreeterImage() *ebiten.Image {
	imgData, err := assets.Assets.ReadFile("image/greeter.png")
	if err != nil {
		return nil
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil
	}
	return ebiten.NewImageFromImage(img)
}
