package objects

type Monster struct {
	Name string
	Life int
	Atk  int
	Def  int
}

func (s *Monster) AttackPlayer(target *Player) {
	target.Life = target.Life + target.Def - s.Atk
}
