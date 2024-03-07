package objects

type Monster struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
	Mdef  int
	// Mdef = Magic Defense
}

func (s *Monster) AttackPlayer(target *Player) {
	if target.Def >= s.Atk {
		target.Life--
	} else {
		target.Life = target.Life + target.Def - s.Atk
	}
}

func (s *Monster) HealSelf() {
	s.Life = s.Life + s.Magic
}

func (s *Monster) IsDead() bool {
	return s.Life <= 0
}
