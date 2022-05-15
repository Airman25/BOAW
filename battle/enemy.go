package battle

var playerKillCount int
var Win int

//calls for all existing enemies
func EnemySkills() {
	if -Skill < 4 {
		Target = 0
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
		for i := 0; i < len(BattleParticipans); i++ {
			BattleParticipans[0].Defense = 0
		}
		playerKillCount = 0
		Skill = 0
	}
}

//calls specific enemy skill
func enemySkill() {
	switch Dict[BattleParticipans[-Skill].Name].Skill {
	case 0:
		grasshopperJump()
	case 1:
		beetleDash()
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
