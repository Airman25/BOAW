package rooms

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

type RoomObject struct {
	X, Y  float64
	W, H  int
	Scale float64
	Image *ebiten.Image
}

const buttonscaling = 1.0

const buttonwidth = 256 * buttonscaling
const buttonheight = 64 * buttonscaling

func Rooms(room, screenWidth, screenHeight int) []RoomObject { //gets what objects need to be loaded in a game "room"
	if room == 0 {
		return []RoomObject{{float64(screenWidth)/2 - buttonwidth/2, float64(screenHeight)/2 - buttonheight*2, buttonwidth, buttonheight, buttonscaling, load.ImageText("button", "Start", buttonwidth/buttonscaling, buttonheight/buttonscaling, "", 0, 0)}}
	}
	return nil
}
