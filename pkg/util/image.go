package util

import (
	"bytes"
	"image"
	"log"
	"strings"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-observer/assets"
)

const (
	roleCount = 17
)

func LoadRoleImages() map[string]*ebiten.Image {
	roleImages := make(map[string]*ebiten.Image)
	for index := range roleCount {
		roleName := strings.ToLower(string(hll.RoleFromInt(index)))
		imgData, err := assets.Assets.ReadFile("roles/" + roleName + ".png")
		if err != nil {
			log.Printf("Error loading role image for %s: %v\n", roleName, err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			log.Printf("Error loading role image for %s: %v\n", roleName, err)
			continue
		}
		roleImages[roleName] = ebiten.NewImageFromImage(img)
	}
	return roleImages
}

func LoadSpawnImages() map[string]*ebiten.Image {
	spawnImages := make(map[string]*ebiten.Image)
	spawnTypes := []string{"garrison", "outpost"}
	for _, spawnType := range spawnTypes {
		imgData, err := assets.Assets.ReadFile("spawns/" + spawnType + ".png")
		if err != nil {
			log.Printf("Error loading role image for %s: %v\n", spawnType, err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			log.Printf("Error loading role image for %s: %v\n", spawnType, err)
			continue
		}
		spawnImages[spawnType] = ebiten.NewImageFromImage(img)
	}
	return spawnImages
}

func LoadMapImage(mapName string) (*ebiten.Image, error) {
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

func LoadGreeterImage() *ebiten.Image {
	imgData, err := assets.Assets.ReadFile("image/greeter.png")
	if err != nil {
		log.Printf("Error loading greeter image: %v\n", err)
		return nil
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Printf("Error decoding greeter image: %v\n", err)
		return nil
	}
	return ebiten.NewImageFromImage(img)
}
