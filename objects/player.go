package objects

import "fmt"

type Player struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
}

// Add a method to the player
func (s *Player) AttackPlayer(target *Player) {
	if !s.IsDead() {
		fmt.Printf("%s attacked %s\n", s.Name, target.Name)
		target.Life = target.Life + target.Def - s.Atk
	} else {
		fmt.Printf("Player Dead - Skip Turn")
	}
}

func (s *Player) AttackMonster(mons *Monster) {
	if !s.IsDead() {
		fmt.Printf("%s attacked %s\n", s.Name, mons.Name)
		mons.Life = mons.Life + mons.Def - s.Atk
	} else {
		fmt.Printf("Player Dead - Skip Turn\n")
	}
}

func (s *Player) Heal(pts *Player) {
	s.Life = s.Life + pts.Magic
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}
