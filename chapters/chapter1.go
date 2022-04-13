package chapters

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const screenWidth = 1280
const screenHeight = 720

func ChapterIAnitmation(animation *int, screen *ebiten.Image, images []*ebiten.Image) {
	if *animation*5-512 < screenWidth/2 { //"Chapter I" text flying from the border to center
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth)-5*float64(*animation)+256, screenHeight/4-screenHeight/8+64)
		screen.DrawImage(images[2], op)
	} else { //as it reaches center it stops flying further
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth/2)-256, screenHeight/4-screenHeight/8+64)
		screen.DrawImage(images[2], op)
	}
	if *animation*5-256 < screenWidth { //first rhinoceros beetle
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth)-5*float64(*animation), screenHeight/4-screenHeight/8)
		screen.DrawImage(images[0], op)
	} else { //second rhinoceros beetle
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(5*float64(*animation)-float64(screenWidth)-512, screenHeight/4+screenHeight/8)
		screen.DrawImage(images[1], op)

		if 5**animation-screenWidth-768 < screenWidth/2 { //"Meadow" text flying from the border to center
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(5*float64(*animation)-float64(screenWidth)-1024, screenHeight/4+screenHeight/8+64)
			screen.DrawImage(images[3], op)
		} else {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(screenWidth/2)-256, screenHeight/4+screenHeight/8+64)
			screen.DrawImage(images[3], op)
		}
	}

	//exit from animation + preporation for the next one
	if *animation > 60*15 { //animation lasts 1*15/load.GameSpeed second
		*animation = 0
	}
}

func Requipment1() []*ebiten.Image {
	return []*ebiten.Image{
		load.ImageEbiten(`\spoilers\` + "chapter_I_0"),
		load.ImageEbiten(`\spoilers\` + "chapter_I_1"),
		load.ImageText("", load.Localisation["chapter1"], 512, 128, "Agreloycalmond", 72, 2),
		load.ImageText("", load.Localisation["chapter1Name"], 512, 128, "Agreloycalmond", 72, 2),
	}
}

func MapLevel() {

}
