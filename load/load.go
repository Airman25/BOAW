package load

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func ImageEbiten(filename string) *ebiten.Image { //Images for ebiten to render
	imageEbiten, _ := imageGet(filename)
	return imageEbiten
}

func ImageNormal(filename string) image.Image { //Images that might be used for other purposes (for example to write text on them before rendering)
	_, imageNormal := imageGet(filename)
	return imageNormal
}

func imageGet(filename string) (*ebiten.Image, image.Image) {
	playImg, image, err := ebitenutil.NewImageFromFile(`src\textures\` + filename + ".png")
	if err != nil {
		log.Fatalf("Load image error: %v", err)
		return nil, nil
	}
	return playImg, image
}
