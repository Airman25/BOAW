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
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*3, buttonwidth, buttonheight, buttonscaling, 1, load.ImageText("button", "Start", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2, buttonwidth, buttonheight, buttonscaling, 2, load.ImageText("button", "Settings", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight, buttonwidth, buttonheight, buttonscaling, 3, load.ImageText("button", "Load", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight) / 2, buttonwidth, buttonheight, buttonscaling, 4, load.ImageText("button", "Authors", buttonwidth, buttonheight, "", 0, 0)},
		{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 + buttonheight, buttonwidth, buttonheight, buttonscaling, 99, load.ImageText("button_exit", "Exit", buttonwidth, buttonheight, "", 0, 0)},
	}
}
