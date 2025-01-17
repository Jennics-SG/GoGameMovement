package app

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// The below line is not a comment it loads in the folder

//go:embed assets/*
var assetLoader embed.FS

// ASSET CACHE CLASS ==========================================================
type AssetCache struct {
	assets []*ebiten.Image
}

// Add asset to asset cache
func (ac *AssetCache) addAsset(path string) {
	a := loadAssetFromFile(path)
	ac.assets = append(ac.assets, a)
}

// ============================================================================

// Loads an asset from a file within the already loaded folder
func loadAssetFromFile(path string) *ebiten.Image {
	f, err := assetLoader.Open(path)

	if err != nil {
		log.Fatalf("ASSET NOT LOADED %v\n", err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("ASSET NOT LOADED %v\n", err)
	}

	return ebiten.NewImageFromImage(img)
}
