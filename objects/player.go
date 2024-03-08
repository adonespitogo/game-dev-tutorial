package objects

import (
	"fmt"
)

type Player struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
}

// Add a method to the player
func (s *Player) AttackPlayer(target *Player) {
	fmt.Printf("%s attacked %s\n", s.Name, target.Name)
	if target.Def >= s.Atk {
		target.Life = target.Life - 1
	} else {
		target.Life = target.Life + target.Def - target.Atk
	}
}

func (s *Player) AttackMonster(mons *Monster) {
	fmt.Printf("%s attacked %s\n", s.Name, mons.Name)
	if mons.Def >= s.Atk {
		mons.Life = mons.Life - 1
	} else {
		mons.Life = mons.Life + mons.Def - s.Atk
	}
}

func (s *Player) Heal(target *Player) {
	fmt.Printf("%s used Heal on %s\n", s.Name, target.Name)
	target.Life = target.Life + s.Magic
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}

func (s *Player) PlayerTurn(m *Monster) {
	if !s.IsDead() {
		s.ChooseMove(m) //Option
	} else {
		fmt.Println("Player Dead - Skip turn")
	}
}

func (s *Player) ChooseMove(m *Monster) {

	fmt.Printf("%s turn\n", s.Name)
	fmt.Printf("1: Attack  2:Heal  3:Skip\n")
	fmt.Printf(">>> ")

	n := "0"
	fmt.Scanln(&n)

	switch n {
	case "1":
		s.AttackMonster(m)
	case "2":
		s.Heal(s)
	case "3":
		fmt.Println("Skip Turn")
	default:
		s.AttackMonster(m)
	}
}

func NewPlayer(n string) *Player {
	fmt.Println("Select Job Class")
	fmt.Print(">>> ")
	fmt.Scanln(&n)
	switch n {
	case "1":
		return &Player{Name: "Mage", Life: 100, Def: 10, Atk: 10, Magic: 30}
	case "2":
		return &Player{Name: "Warrior", Life: 100, Def: 15, Atk: 20, Magic: 5}
	case "3":
		return &Player{Name: "Archer", Life: 100, Def: 10, Atk: 30, Magic: 5}
	default:
		return &Player{Name: "Jobless", Life: 100, Def: 10, Atk: 10, Magic: 5}
	}
}
