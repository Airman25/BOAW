package battle

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

//our first and currently only hero
func Hero1() []BattleObject {
	return []BattleObject{
		{screenWidth / 8, screenHeight / 2, 2, 0, "hero", 50, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "player_right"), load.ImageEbiten(`\spoilers\` + "player_right_2"), load.ImageEbiten(`\spoilers\` + "player_left"), load.ImageEbiten(`\spoilers\` + "player_left_2"), load.ImageEbiten(`\spoilers\` + "player_shield")}},
	}
}

//calls specific hero's skill
func HeroSkills() {
	startY = screenHeight / 2
	if Skill == 1 {
		basicAttack()
	}
	if Skill == 2 {
		basicBlock()
	}
}
