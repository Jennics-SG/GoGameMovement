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

// ASSET STRUCT ===============================================================
type Asset struct {
	id    string
	image *ebiten.Image
}

func NewAsset(id string, image *ebiten.Image) *Asset {
	return &Asset{id: id, image: image}
}

// ASSET CACHE CLASS ==========================================================
type AssetCache struct {
	assets []*Asset
}

// Add asset to asset cache
func (ac *AssetCache) addAsset(path string, id string) {
	image := loadAssetFromFile(path)

	a := NewAsset(id, image)

	ac.assets = append(ac.assets, a)
}

// Search through cache for asset with ID matching
func (ac *AssetCache) find(id string) *ebiten.Image {
	for _, asset := range ac.assets {
		if asset.id == id {
			return asset.image
		}
	}
	return nil
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
