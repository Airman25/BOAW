package battle

func elementalAttack(owner string) {
	BattleParticipans[Target].Health -= int(float64(Dict[owner].Damage) * (1 - Dict[BattleParticipans[Target].Name].Resistance[Dict[owner].element]) * (1 - BattleParticipans[Target].Defense))
	damageDealt = true
}
