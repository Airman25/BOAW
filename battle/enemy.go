package battle

var playerWin int
var Win int

//calss all existing enemies
func EnemySkills(everyone []BattleObject) {
	if len(everyone) > -Skill {
		if everyone[-Skill].Health > 0 { // Skip enemies with no HP left
			enemySkill(everyone)
		} else {
			playerWin++
			Skill--
		}
	} else {
		if -Skill-1 == playerWin {
			Win = -Win
		}
		playerWin = 0
		Skill = 0
	}
}

//calls specific enemy skill
func enemySkill(everyone []BattleObject) {
	if dict[everyone[-Skill].Name].Skills[0] == 0 {
		grasshopperJump(everyone)
	}
}
