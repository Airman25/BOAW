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
	Defense    float64
	ImageInUse int
	Images     []*ebiten.Image
}

var BattleParticipans []BattleObject

//data about which enemies exist in battle
func Index(index int) []BattleObject {
	Win = -index
	switch index {
	case 1:
		return enemys("grasshopper", "", "")
	case 2:
		return enemys("grasshopper", "grasshopper", "")
	case 3:
		return enemys("rhinoceros_beetle", "rhinoceros_beetle", "")
	case 4:
		return enemys("grasshopper", "rhinoceros_beetle", "")
	}
	return nil
}
