package objects

type Monster struct {
	Name string
	Life int
	Atk  int
	Def  int
}

func (mons *Monster) AttackPlayer(target *Player) {
	target.Life = target.Life + target.Def - mons.Atk
}
