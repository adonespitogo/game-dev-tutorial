package objects

import "fmt"

type Player struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
	MDef  int
	Dex   int
	//Mdef = Magic defense
}

// Add a method to the player
func (s *Player) AttackPlayer(target *Player) {

	fmt.Printf("%s attacked %s\n", s.Name, target.Name)
	target.Life = target.Life + target.Def - s.Atk
}

func (s *Player) AttackMonster(mons *Monster) {
	fmt.Printf("%s attacked %s\n", s.Name, mons.Name)
	mons.Life = mons.Life + mons.Def - s.Atk
}

func (s *Player) HealSelf() {
	fmt.Printf("%s used Heal on self\n", s.Name)
	s.Life = s.Life + s.Magic
}

func (s *Player) HealOther(pts *Player) {
	fmt.Printf("%s used Heal on %s", s.Name, pts.Name)
	s.Life = s.Life + pts.Magic
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}

// Matk = Magic Attack
func (s *Player) MAtkPlayer(target *Player) {
	target.Life = target.Life + target.MDef - s.Magic
}

func (s *Player) MAtkMons(mons *Monster) {
	mons.Life = mons.Life + mons.Def - s.Magic
}

func (s *Player) PlayerTurn(m *Monster) {
	if !s.IsDead() {
		n := "0"
		fmt.Scanln(&n)
		s.ChooseMove(n, m) //Option
	} else {
		fmt.Println("Player Dead - Skip turn")
	}
}

func (s *Player) ChooseMove(n string, m *Monster) {

	switch n {
	case "1":
		s.AttackMonster(m)
	case "2":
		s.HealSelf()
	case "3":
		fmt.Println("Skip Turn")
	default:
		s.AttackMonster(m)
	}
}
