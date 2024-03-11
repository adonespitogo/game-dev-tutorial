package objects

import (
	"fmt"
	"time"
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
	DialogSpacer(1)
	DotLoad(5)
	fmt.Printf("%s attacked %s\n", s.Name, target.Name)
	DialogSpacer(1)
	if target.Def >= s.Atk {
		target.Life = target.Life - 1
	} else {
		target.Life = target.Life + target.Def - target.Atk
	}
}

func (s *Player) AttackMonster(mons *Monster) {
	DialogSpacer(1)
	fmt.Printf("%s launches an attack on %s ", s.Name, mons.Name)
	DotLoad(5)
	fmt.Printf("%s attacked %s\n", s.Name, mons.Name)
	time.Sleep(1 * time.Second)
	DialogSpacer(1)
	if mons.Def >= s.Atk {
		mons.Life = mons.Life - 1
	} else {
		mons.Life = mons.Life + mons.Def - s.Atk
	}
	DialogSpacer(1)
	fmt.Printf("%s HP went down to %d\n", mons.Name, mons.Life)
	DialogSpacer(1)
}

func (s *Player) Heal(target *Player) {
	DialogSpacer(1)
	fmt.Printf("%s is casting Heal", s.Name)
	DotLoad(5)
	fmt.Printf("%s used Heal on %s\n", s.Name, target.Name)
	time.Sleep(1 * time.Second)
	DialogSpacer(1)
	target.Life = target.Life + s.Magic
}

func (s *Player) IsDead() bool {
	return s.Life <= 0
}

func (s *Player) PlayerTurn(m *Monster) {
	if !s.IsDead() {
		s.ChooseMove(m) //Option
	} else {
		DialogSpacer(1)
		fmt.Println("Player Dead - Skip turn")
		DialogSpacer(1)

	}
}

func (s *Player) ChooseMove(m *Monster) {

	DialogSpacer(1)
	fmt.Printf(">> %s turn <<\n", s.Name)

	DialogSpacer(1)
	PlayerStatusBar(s, m)
	DialogSpacer(1)
	SkillBox(s)

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

func NewPlayer(j string) *Player {

	DialogSpacer(1)
	n := "No name"
	switch j {
	case "1": //mage
		fmt.Println("Enter Mage Name")
		fmt.Print(">>> ")
		fmt.Scanln(&n)
		DotLoad(3)
		fmt.Printf("Player %s has been created", n)
		DialogSpacer(1)
		return &Player{Name: n, Life: 100, Def: 10, Atk: 10, Magic: 30}
	case "2": //warrior
		fmt.Println("Enter Warrior Name")
		fmt.Print(">>> ")
		fmt.Scanln(&n)
		DotLoad(3)
		fmt.Printf("Player %s has been created", n)
		DialogSpacer(1)
		return &Player{Name: n, Life: 150, Def: 15, Atk: 20, Magic: 5}
	case "3": //archer
		fmt.Println("Enter Archer Name")
		fmt.Print(">>> ")
		fmt.Scanln(&n)
		DotLoad(3)
		fmt.Printf("Player %s has been created", n)
		DialogSpacer(1)
		return &Player{Name: n, Life: 120, Def: 10, Atk: 30, Magic: 5}
	default: //jobless
		return &Player{Name: n, Life: 90, Def: 10, Atk: 10, Magic: 5}
	}

}

func SkillBox(s *Player) {
	fmt.Printf("<>===================<>\n")
	fmt.Printf("||  1 - Attack (%d)  ||\n", s.Atk)
	if s.Magic > 9 {
		fmt.Printf("||  2 - Heal (%d)    ||\n", s.Magic)
	} else {
		fmt.Printf("||  2 - Heal (%d)     ||\n", s.Magic)
	}
	fmt.Printf("||  3 - Skip Turn    ||\n")
	fmt.Printf("<>===================<>\n")
	DialogSpacer(1)
	fmt.Printf(">>> ")
}

func PlayerStatusBar(s *Player, target *Monster) {
	fmt.Printf("%s		|	%s\n", s.Name, target.Name)
	fmt.Printf("HP:%d	DEF:%d  |	HP:%d	DEF:%d  \n", s.Life, s.Def, target.Life, target.Def)

	fmt.Printf("%s\n", s.Name)
	fmt.Printf("HP  :%d   Def :%d\n", s.Life, s.Def)
	fmt.Printf("Atk :%d   MAG :%d", s.Atk, s.Magic)
}
