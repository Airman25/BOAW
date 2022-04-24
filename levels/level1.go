package levels

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimatedObject struct {
	X, Y       float64
	W, H       float64
	Scale      float64
	Data       int
	ImageInUse int
	Images     []*ebiten.Image
}

var Player = AnimatedObject{250, 250, 0, 0, 2, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "player_right"), load.ImageEbiten(`\spoilers\` + "player_left"), load.ImageEbiten(`\spoilers\` + "player_up"), load.ImageEbiten(`\spoilers\` + "player_down"), load.ImageEbiten(`\spoilers\` + "player_right_2"), load.ImageEbiten(`\spoilers\` + "player_left_2"), load.ImageEbiten(`\spoilers\` + "player_up_2"), load.ImageEbiten(`\spoilers\` + "player_down_2")}}

func Location1() []AnimatedObject {
	return []AnimatedObject{
		{500, 250, 64, 64, 1, 1, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "grasshopper")}},
	}
}
