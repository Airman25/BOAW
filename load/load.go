package load

import (
	"image"
	_ "image/png"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type moreData struct {
	fs    float64
	font  string
	color int
}

func ImageEbiten(filename string) *ebiten.Image { //Images for ebiten to render
	imageEbiten, _ := imageGet(filename)
	return imageEbiten
}

func ImageNormal(filename string) image.Image { //Images that might be used for other purposes (for example to write text on them before rendering)
	_, imageNormal := imageGet(filename)
	return imageNormal
}

//gets the image from src folder
func imageGet(filename string) (*ebiten.Image, image.Image) {
	ebitenImg, img, err := ebitenutil.NewImageFromFile(`src\textures\` + filename + ".png")
	if err != nil {
		log.Fatalf("Load image error: %v", err)
		return nil, nil
	}
	return ebitenImg, img
}

func ImageText(filename, text string, width, height int, font string, fs float64, color int) *ebiten.Image {
	return ebiten.NewImageFromImage(imageSetText(ImageNormal(filename), text, width, height, font, fs, color))
}

//simply adds text to image
func imageSetText(img image.Image, textonimage string, width int, height int, font string, fs float64, color int) image.Image {
	if font == "" { //default case for font and fontsize
		font = "arial"
	}
	if fs == 0 {
		fs = 24
	}
	dc := gg.NewContext(width, height)
	if err := dc.LoadFontFace(`src\fonts\`+font+".ttf", fs); err != nil {
		panic(err)
	}
	if img != nil {
		dc.DrawImage(img, 0, 0)
	}
	if color == 0 { // 0 = black, 1 = white
		dc.SetRGB(0, 0, 0)
	} else {
		dc.SetRGB(1, 1, 1)
	}
	dc.DrawStringAnchored(textonimage, float64(width)/2, float64(height)/2, 0.5, 0.5)
	dc.Clip()
	return dc.Image()
}
