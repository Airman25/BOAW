package battle

var playerKillCount int
var Win int

//calls for all existing enemies
func EnemySkills() {
	if -Skill < 4 {
		if Alive(-Skill) {
			enemySkill()
		} else {
			playerKillCount++
			Skill--
		}
	} else {
		if BattleParticipans[0].Health > 0 {
			if -Skill-1 == playerKillCount {
				Win = -Win
			}
		} else {
			Win = 404
		}
		playerKillCount = 0
		Skill = 0
	}
}

//calls specific enemy skill
func enemySkill() {
	if dict[BattleParticipans[-Skill].Name].Skills[0] == 0 {
		grasshopperJump()
	}
}

func Alive(specific int) bool { //checks if specified enemy is alive
	if len(BattleParticipans) > specific {
		if BattleParticipans[specific].Health > 0 { // Skip enemies with no HP left
			switch specific {
			case 1:
				startY = screenHeight / 2
			case 2:
				startY = screenHeight/2 - screenHeight/4
			case 3:
				startY = screenHeight/2 + screenHeight/4
			}
			return true
		}
	}
	return false
}
