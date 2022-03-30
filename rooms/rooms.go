package rooms

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

type RoomObject struct {
	X, Y  int
	Image *ebiten.Image
}

func Rooms(room int) []RoomObject { //gets what objects need to be loaded in a game "room"
	if room == 0 {
		return []RoomObject{{0, 0, load.ImageEbiten("button0")}}
	}
	return nil
}
