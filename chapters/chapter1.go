package chapters

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const screenWidth = 1280
const screenHeight = 720
const moveSpeed = 7

func ChapterIAnitmation(animation *int, screen *ebiten.Image, images []*ebiten.Image) {
	if *animation*moveSpeed-512 < screenWidth/2 { //"Chapter I" text flying from the border to center
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth)-moveSpeed*float64(*animation)+256, screenHeight/4-screenHeight/8+64)
		screen.DrawImage(images[2], op)
	} else { //as it reaches center it stops flying further
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth/2)-256, screenHeight/4-screenHeight/8+64)
		screen.DrawImage(images[2], op)
	}
	if *animation*moveSpeed-256 < screenWidth { //first rhinoceros beetle
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth)-moveSpeed*float64(*animation), screenHeight/4-screenHeight/8)
		screen.DrawImage(images[0], op)
	} else { //second rhinoceros beetle
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(moveSpeed*float64(*animation)-float64(screenWidth)-512, screenHeight/4+screenHeight/8)
		screen.DrawImage(images[1], op)

		if *animation*moveSpeed-screenWidth-768 < screenWidth/2 { //"Meadow" text flying from the border to center
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(moveSpeed*float64(*animation)-float64(screenWidth)-1024, screenHeight/4+screenHeight/8+64)
			screen.DrawImage(images[3], op)
		} else {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(screenWidth/2)-256, screenHeight/4+screenHeight/8+64)
			screen.DrawImage(images[3], op)
		}
	}

	//exit from animation + preporation for the next one
	if *animation > 60*10 { //animation lasts 1*10/load.GameSpeed second
		*animation = 0
	}
}

func Requipment1() []*ebiten.Image {
	if load.Localisation["chapter1Name"] == "Meadow" { //if language == english
		return []*ebiten.Image{
			load.ImageEbiten(`\spoilers\` + "chapter_I_0"),
			load.ImageEbiten(`\spoilers\` + "chapter_I_1"),
			load.ImageText("", load.Localisation["chapter1"], 512, 128, "Agreloycalmond", 72, 2),
			load.ImageText("", load.Localisation["chapter1Name"], 512, 128, "Agreloycalmond", 72, 2),
		}
	}
	return []*ebiten.Image{ //can't use Agreloycalmond font with other languages
		load.ImageEbiten(`\spoilers\` + "chapter_I_0"),
		load.ImageEbiten(`\spoilers\` + "chapter_I_1"),
		load.ImageText("", load.Localisation["chapter1"], 512, 128, "", 72, 2),
		load.ImageText("", load.Localisation["chapter1Name"], 512, 128, "", 60, 2),
	}
}
