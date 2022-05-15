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

var startY = screenHeight / 2.0
var backX float64

var Target int

func basicAttack() {
	animateHero += load.GameSpeed
	if screenWidth/8+float64(animateHero)*movementSpeed < screenWidth/2+screenWidth/4 {
		if BattleParticipans[0].Y > BattleParticipans[Target].Y {
			BattleParticipans[0].Y -= movementSpeed
		} else if BattleParticipans[0].Y < BattleParticipans[Target].Y {
			BattleParticipans[0].Y += movementSpeed
		} else if backX == 0 {
			backX = BattleParticipans[0].X
		}
		if animateHero%6 == 0 {
			if BattleParticipans[0].ImageInUse == 0 {
				BattleParticipans[0].ImageInUse = 1
			} else {
				BattleParticipans[0].ImageInUse = 0
			}
		}
		BattleParticipans[0].X = screenWidth/8 + float64(animateHero)*movementSpeed
	} else if BattleParticipans[0].X > screenWidth/8 {
		if !damageDealt {
			BattleParticipans[0].ImageInUse = 2
			elementalAttack("hero")
		}
		if BattleParticipans[0].X <= backX {
			if BattleParticipans[0].Y > startY {
				BattleParticipans[0].Y -= movementSpeed
			} else if BattleParticipans[0].Y < startY {
				BattleParticipans[0].Y += movementSpeed
			}
		}
		if animateHero%6 == 0 {
			if BattleParticipans[0].ImageInUse == 2 {
				BattleParticipans[0].ImageInUse = 3
			} else {
				BattleParticipans[0].ImageInUse = 2
			}
		}
		BattleParticipans[0].X = screenWidth + screenWidth/2 - float64(animateHero)*movementSpeed - screenWidth/8
	} else {
		BattleParticipans[0].ImageInUse = 0
		Skill = -1
		animateHero = 0
		backX = 0
		damageDealt = false
	}
}

func basicBlock() {
	animateHero += load.GameSpeed
	if animateHero < 60*1.5 {
		BattleParticipans[0].ImageInUse = 4
	} else {
		BattleParticipans[0].ImageInUse = 0
		Skill = -1
		animateHero = 0
		BattleParticipans[0].Defense = 0.5 //reduce damage taken by 50%
	}
}

func grasshopperJump() {
	animateEnemy += load.GameSpeed
	if screenWidth/2+screenWidth/4-float64(animateEnemy)*movementSpeed > screenWidth/8+heroWidth {
		if BattleParticipans[-Skill].Y > BattleParticipans[0].Y {
			BattleParticipans[-Skill].Y -= movementSpeed
		} else if BattleParticipans[-Skill].Y < BattleParticipans[0].Y {
			BattleParticipans[-Skill].Y += movementSpeed
		} else if backX == 0 {
			backX = BattleParticipans[-Skill].X
		}
		BattleParticipans[-Skill].X = screenWidth/2 + screenWidth/4 - float64(animateEnemy)*movementSpeed
	} else if math.Round(BattleParticipans[-Skill].Rotation*100) != 628 { //starts rotating at 0 and stops on math.Pi*2
		if !damageDealt {
			elementalAttack(BattleParticipans[-Skill].Name)
		}
		BattleParticipans[-Skill].Rotation += math.Pi / (18 / float64(load.GameSpeed))
	} else if BattleParticipans[-Skill].X < screenWidth/2+screenWidth/4 {
		if BattleParticipans[-Skill].X >= backX {
			if BattleParticipans[-Skill].Y > startY {
				BattleParticipans[-Skill].Y -= movementSpeed
			} else if BattleParticipans[-Skill].Y < startY {
				BattleParticipans[-Skill].Y += movementSpeed
			}
		}
		BattleParticipans[-Skill].X += float64(load.GameSpeed) * movementSpeed
	} else {
		BattleParticipans[-Skill].Rotation = 0 //reset rotation
		Skill--
		animateEnemy = 0
		backX = 0
		damageDealt = false
	}
}

func beetleDash() {
	animateEnemy += load.GameSpeed
	if screenWidth/2+screenWidth/4-float64(animateEnemy)*movementSpeed > screenWidth/8 {
		if BattleParticipans[-Skill].Y > BattleParticipans[0].Y {
			BattleParticipans[-Skill].Y -= movementSpeed
		} else if BattleParticipans[-Skill].Y < BattleParticipans[0].Y {
			BattleParticipans[-Skill].Y += movementSpeed
		} else if backX == 0 {
			backX = BattleParticipans[-Skill].X
		} else {
			animateEnemy += load.GameSpeed * 2
		}
		BattleParticipans[-Skill].X = screenWidth/2 + screenWidth/4 - float64(animateEnemy)*movementSpeed
	} else if BattleParticipans[-Skill].X < screenWidth/2+screenWidth/4 {
		if !damageDealt {
			elementalAttack(BattleParticipans[-Skill].Name)
		}
		if BattleParticipans[-Skill].X >= backX {
			if BattleParticipans[-Skill].Y > startY {
				BattleParticipans[-Skill].Y -= movementSpeed
			} else if BattleParticipans[-Skill].Y < startY {
				BattleParticipans[-Skill].Y += movementSpeed
			}
		}
		BattleParticipans[-Skill].X += float64(load.GameSpeed) * movementSpeed
	} else {
		Skill--
		animateEnemy = 0
		backX = 0
		damageDealt = false
	}
}
