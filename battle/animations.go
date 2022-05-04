package battle

import (
	"math"

	"github.com/Airman25/BOAW/load"
)

var Skill int
var animateHero int
var animateEnemy int
var damageDealt bool

const movementSpeed = 12
const heroWidth = 128

const startY = screenHeight / 2

var Target int

func basicAttack(everyone []BattleObject) {
	animateHero += load.GameSpeed
	if screenWidth/8+float64(animateHero)*movementSpeed < screenWidth/2+screenWidth/4 {
		if everyone[0].Y > everyone[Target].Y {
			everyone[0].Y -= movementSpeed
		} else if everyone[0].Y < everyone[Target].Y {
			everyone[0].Y += movementSpeed
		}
		if animateHero%6 == 0 {
			if everyone[0].ImageInUse == 0 {
				everyone[0].ImageInUse = 1
			} else {
				everyone[0].ImageInUse = 0
			}
		}
		everyone[0].X = screenWidth/8 + float64(animateHero)*movementSpeed
	} else if everyone[0].X > screenWidth/8 {
		if !damageDealt {
			everyone[0].ImageInUse = 2
			everyone[Target].Health -= 10
			damageDealt = true
		}
		if everyone[0].Y > startY {
			everyone[0].Y -= movementSpeed
		} else if everyone[0].Y < startY {
			everyone[0].Y += movementSpeed
		}
		if animateHero%6 == 0 {
			if everyone[0].ImageInUse == 2 {
				everyone[0].ImageInUse = 3
			} else {
				everyone[0].ImageInUse = 2
			}
		}
		everyone[0].X = screenWidth + screenWidth/2 - float64(animateHero)*movementSpeed - screenWidth/8
	} else {
		everyone[0].ImageInUse = 0
		Skill = -1
		animateHero = 0
		damageDealt = false
	}
}

func grasshopperJump(everyone []BattleObject) {
	animateEnemy += load.GameSpeed
	if screenWidth/2+screenWidth/4-float64(animateEnemy)*movementSpeed > screenWidth/8+heroWidth {
		everyone[-Skill].X = screenWidth/2 + screenWidth/4 - float64(animateEnemy)*movementSpeed
	} else if math.Round(everyone[-Skill].Rotation*100) != 628 { //starts rotating at 0 and stops on math.Pi*2
		if !damageDealt {
			everyone[0].Health -= dict[everyone[-Skill].Name].Damage
			damageDealt = true
		}
		everyone[-Skill].Rotation += math.Pi / (18 / float64(load.GameSpeed))
	} else if everyone[-Skill].X < screenWidth/2+screenWidth/4 {
		everyone[-Skill].X += float64(load.GameSpeed) * movementSpeed
	} else {
		everyone[-Skill].Rotation = 0 //reset rotation
		Skill--
		animateEnemy = 0
		damageDealt = false
	}
}
