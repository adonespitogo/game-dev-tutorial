package objects

type Player struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
}

// Add a method to the player
func (s *Player) AttackPlayer(target *Player) {
	target.Life = target.Life + target.Def - s.Atk
}

func (s *Player) AttackMonster(mons *Monster) {
	mons.Life = mons.Life + mons.Def - s.Atk
}

func (s *Player) Heal(pts *Player) {
	s.Life = s.Life + pts.Magic
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}
