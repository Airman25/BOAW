package battle

import (
	"math"

	"github.com/Airman25/BOAW/load"
)

var Skill int
var animateHero int
var animateEnemy int
var dealDamage bool

const movementSpeed = 12
const heroWidth = 128

func basicAttack(everyone []BattleObject) {
	animateHero += load.GameSpeed
	if screenWidth/8+float64(animateHero)*movementSpeed < screenWidth/2+screenWidth/4 {
		if animateHero%6 == 0 {
			if everyone[0].ImageInUse == 0 {
				everyone[0].ImageInUse = 1
			} else {
				everyone[0].ImageInUse = 0
			}
		}
		everyone[0].X = screenWidth/8 + float64(animateHero)*movementSpeed
	} else if everyone[0].X > screenWidth/8 {
		if animateHero%6 == 0 {
			if everyone[0].ImageInUse == 2 {
				everyone[0].ImageInUse = 3
			} else {
				everyone[0].ImageInUse = 2
			}
		}
		everyone[0].X = screenWidth + screenWidth/2 - float64(animateHero)*movementSpeed - screenWidth/8
	}
	if !dealDamage && animateHero > 60*1.25 {
		everyone[0].ImageInUse = 2
		everyone[len(everyone)-1].Health -= 10
		dealDamage = true
	}

	if animateHero > 60*2.5 {
		everyone[0].ImageInUse = 0
		Skill = -1
		animateHero = 0
		dealDamage = false
	}

}

func grasshopperJump(everyone []BattleObject) {
	animateEnemy += load.GameSpeed
	if screenWidth/2+screenWidth/4-float64(animateEnemy)*movementSpeed > screenWidth/8+heroWidth {
		everyone[-Skill].X = screenWidth/2 + screenWidth/4 - float64(animateEnemy)*movementSpeed
	} else if animateEnemy <= 60*1.5+1 {
		everyone[-Skill].Rotation += math.Pi / 18 * float64(load.GameSpeed)
	} else if everyone[-Skill].X < screenWidth/2+screenWidth/4 {
		everyone[-Skill].X += float64(load.GameSpeed) * movementSpeed
	}
	if !dealDamage && animateEnemy > 60*1.5 {
		everyone[0].Health -= 10
		dealDamage = true
	}
	if animateEnemy > 60*2.5 {
		Skill--
		animateEnemy = 0
		dealDamage = false
	}
}
