package objects

type Monster struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
}

func (s *Monster) AttackPlayer(target *Player) {
	target.Life = target.Life + target.Def - s.Atk
}

func (s *Monster) Heal(pts *Monster) {
	s.Life = s.Life + pts.Magic
}

func (s *Monster) IsDead() bool {
	return s.Life <= 0
}
