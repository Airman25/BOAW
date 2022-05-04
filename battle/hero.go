package battle

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

//our first and currently only hero
func Hero1() []BattleObject {
	return []BattleObject{
		{screenWidth / 8, screenHeight / 2, 2, 0, "hero1", 50, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + "player_right"), load.ImageEbiten(`\spoilers\` + "player_right_2"), load.ImageEbiten(`\spoilers\` + "player_left"), load.ImageEbiten(`\spoilers\` + "player_left_2")}},
	}
}

//calls specific hero's skill
func HeroSkills(everyone []BattleObject) {
	if Skill == 1 {
		basicAttack(everyone)
	}
}
