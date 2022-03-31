package rooms

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const buttonscaling = 1.0

const buttonwidth = 256
const buttonheight = 64

func Start() {

}

type RoomObject struct {
	X, Y     float64
	W, H     int
	Scale    float64
	Function int
	Image    *ebiten.Image
}

func GetRoom1(screenWidth, screenHeight int) []RoomObject {
	return []RoomObject{{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2, buttonwidth, buttonheight, buttonscaling, 1, load.ImageText("button", "Start", buttonwidth, buttonheight, "", 0, 0)}}
}
