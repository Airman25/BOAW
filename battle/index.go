package battle

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BattleObject struct {
	X, Y       float64
	Scale      float64
	Rotation   float64
	Name       string
	Health     int
	ImageInUse int
	Images     []*ebiten.Image
}

//data about which enemies exist in battle
func Index(index int) []BattleObject {
	Win = -index
	switch index {
	case 1:
		return enemys("grasshopper", "", "")
	case 2:
		return enemys("grasshopper", "grasshopper", "")
	}
	return nil
}
