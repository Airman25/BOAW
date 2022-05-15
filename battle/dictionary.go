package battle

import (
	"github.com/Airman25/BOAW/load"
	"github.com/hajimehoshi/ebiten/v2"
)

const screenWidth = 1280
const screenHeight = 720

/*
Resistence:
0	Fire
1	Water
2	Earth
3	Wind
*/

type enemy struct {
	MaxHP      int
	Damage     int
	element    int
	Skill      int
	Resistance [4]float64 //elemental resistance in %
}

var Dict = make(map[string]*enemy)

func Initialise() {
	addenemy("hero", 100, 10, 2, 0, [4]float64{-0.25, 0.25, 0.5, 0.25}) //initial stats
	addenemy("grasshopper", 20, 10, 2, 0, [4]float64{-0.25, -0.25, 0.5, 0.5})
	addenemy("rhinoceros_beetle", 20, 10, 2, 1, [4]float64{-0.25, -0.25, 0.5, 0.5})
	addenemy("spider", 0, 0, 2, 0, [4]float64{-0.25, -0.25, 0.5, 0.5})
	addenemy("bee", 0, 0, 2, 0, [4]float64{-0.25, -0.25, 0.5, -0.25})
}

//adds enemy to dictionary
func addenemy(name string, hp int, damage int, el int, skill int, res [4]float64) {
	Dict[name] = &enemy{hp, damage, el, skill, res}
}

//enemies placement in battle
func enemys(enemy1 string, enemy2 string, enemy3 string) []BattleObject {
	enemy := []BattleObject{{screenWidth/2 + screenWidth/4, screenHeight / 2, 2, 0, enemy1, Dict[enemy1].MaxHP, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy1)}}}
	if enemy2 != "" {
		enemy = append(enemy, BattleObject{screenWidth/2 + screenWidth/4, screenHeight/2 - screenHeight/4, 2, 0, enemy2, Dict[enemy2].MaxHP, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy2)}})
		if enemy3 != "" {
			enemy = append(enemy, BattleObject{screenWidth/2 + screenWidth/4, screenHeight/2 + screenHeight/4, 2, 0, enemy3, Dict[enemy3].MaxHP, 0, 0, []*ebiten.Image{load.ImageEbiten(`\spoilers\` + enemy3)}})
		}
	}
	return enemy
}
