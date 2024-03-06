package objects

type Player struct {
	Name string
	Life int
	Atk  int
	Def  int
}

// Add a method to the player
func (s *Player) AttackPlayer(target *Player) {
	target.Life = target.Life + target.Def - s.Atk
}

func (s *Player) AttackMonster(mons *Monster) {
	mons.Life = mons.Life - s.Atk
}

func (s *Player) Heal(pts int) {
	s.Life = s.Life + pts
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}
