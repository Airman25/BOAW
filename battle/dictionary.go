package battle

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const screenWidth = 1280
const screenHeight = 720

type enemy struct {
	MaxHP  int
	Damage int
	Skills []int
}

var dict = make(map[string]*enemy)

func Initialise() {
	addenemy("grasshopper", 40, 10, []int{0})
	addenemy("rhinoceros_beetle", 0, 0, []int{0})
	addenemy("spider", 0, 0, []int{0})
	addenemy("bee", 0, 0, []int{0})
}

//adds enemy to dictionary
func addenemy(name string, hp int, damage int, skills []int) {
	dict[name] = &enemy{hp, damage, skills}
}

//enemies placement in battle
func enemys(enemy1 string, enemy2 string, enemy3 string) []BattleObject {
	if enemy2 != "" {
		if enemy3 != "" {
			return []BattleObject{
				{screenWidth/2 + screenWidth/4, screenHeight / 2, 2, 0, enemy1, dict[enemy1].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy1)}},
				{screenWidth/2 + screenWidth/4, screenHeight/2 - screenHeight/4, 2, 0, enemy2, dict[enemy2].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy2)}},
				{screenWidth/2 + screenWidth/4, screenHeight/2 + screenHeight/4, 2, 0, enemy3, dict[enemy3].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy3)}},
			}
		}
		return []BattleObject{
			{screenWidth/2 + screenWidth/4, screenHeight / 2, 2, 0, enemy1, dict[enemy1].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy1)}},
			{screenWidth/2 + screenWidth/4, screenHeight/2 - screenHeight/4, 2, 0, enemy2, dict[enemy2].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy2)}},
		}
	}
	return []BattleObject{
		{screenWidth/2 + screenWidth/4, screenHeight / 2, 2, 0, enemy1, dict[enemy1].MaxHP, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy1)}},
	}
}
