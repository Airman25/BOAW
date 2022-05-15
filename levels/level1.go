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

type Location struct{}

var Player = AnimatedObject{250, 250, 0, 0, 2, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "player_right"), load.ImageEbiten(`\spoilers\` + "player_left"), load.ImageEbiten(`\spoilers\` + "player_up"), load.ImageEbiten(`\spoilers\` + "player_down"), load.ImageEbiten(`\spoilers\` + "player_right_2"), load.ImageEbiten(`\spoilers\` + "player_left_2"), load.ImageEbiten(`\spoilers\` + "player_up_2"), load.ImageEbiten(`\spoilers\` + "player_down_2")}}

var DefeatedEnemies = make([]int, 4)

func (L Location) LocationX11Y11() []AnimatedObject {
	var enemies []AnimatedObject
	//walls
	enemies = append(enemies, AnimatedObject{0, 0, 1280, 16, 1, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "wall_1")}})
	//check alive enemies and adds them to level
	if DefeatedEnemies[0] == 0 {
		enemies = append(enemies, AnimatedObject{500, 250, 64, 64, 1, 1, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "grasshopper")}})
	}
	return enemies
}

func (L Location) LocationX12Y11() []AnimatedObject {
	var enemies []AnimatedObject
	//walls
	enemies = append(enemies, AnimatedObject{0, 0, 1280, 16, 1, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "wall_1")}})
	//check alive enemies and adds them to level
	if DefeatedEnemies[1] == 0 {
		enemies = append(enemies, AnimatedObject{1200, 300, 64, 64, 1, 2, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "grasshopper")}})
	}
	if DefeatedEnemies[2] == 0 {
		enemies = append(enemies, AnimatedObject{600, 100, 64, 64, 1, 3, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "rhinoceros_beetle")}})
	}
	return enemies
}

func (L Location) LocationX13Y11() []AnimatedObject {
	var enemies []AnimatedObject
	//walls
	enemies = append(enemies, AnimatedObject{0, 0, 1280, 16, 1, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "wall_1")}})
	//check alive enemies and adds them to level
	if DefeatedEnemies[3] == 0 {
		enemies = append(enemies, AnimatedObject{600, 100, 64, 64, 1, 4, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "rhinoceros_beetle")}})
	}
	return enemies
}
