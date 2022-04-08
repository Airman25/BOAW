package rooms

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const buttonscaling = 1.0

const buttonwidth = 256
const buttonheight = 64

type RoomObject struct {
	X, Y     float64
	W, H     int
	Scale    float64
	Function int
	Image    *ebiten.Image
}

//start menu
func GetRoom0(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{
		//buttons
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight - 12, buttonwidth, buttonheight, buttonscaling, 1, load.ImageText("button", load.Localisation["buttonStart"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - 10, buttonwidth, buttonheight, buttonscaling, 2, load.ImageText("button", load.Localisation["buttonContinue"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight - 8, buttonwidth, buttonheight, buttonscaling, 3, load.ImageText("button", load.Localisation["buttonOptions"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*2 - 6, buttonwidth, buttonheight, buttonscaling, 4, load.ImageText("button", load.Localisation["buttonMedals"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*3 - 4, buttonwidth, buttonheight, buttonscaling, 5, load.ImageText("button", load.Localisation["buttonCredits"], buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight*4 - 2, buttonwidth, buttonheight, buttonscaling, 100, load.ImageText("button_exit", load.Localisation["buttonExit"], buttonwidth, buttonheight, "", 0, 0)},
		//text
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight*5, 0, 0, buttonscaling, -1, load.ImageText("", "B   AW", buttonwidth*2, buttonheight*2, "Agreloycalmond", 144, 2)},
		{float64(screenWidth)/2 - buttonwidth - buttonwidth/4 + buttonwidth/32, float64(screenHeight)/2 - buttonheight*5, 0, 0, buttonscaling, -1, load.ImageText("", "O", buttonwidth*2, buttonheight*2, "Agreloycalmond", 168, 3)},
		{float64(screenWidth)/2 - buttonwidth, float64(screenHeight)/2 - buttonheight*4 + buttonheight/2, 0, 0, buttonscaling, -1, load.ImageText("", load.Localisation["descBOAW"], buttonwidth*2, buttonheight*2, "", 36, 2)},
	}
}

func DefaultBackground() []*ebiten.Image {
	return []*ebiten.Image{} //load.ImageEbiten("Back0"), load.ImageEbiten("Back1")
}
