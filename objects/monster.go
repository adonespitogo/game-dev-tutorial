package objects

type Monster struct {
	Name  string
	Life  int
	Atk   int
	Def   int
	Magic int
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

func NewGoblin(lvl int) *Monster {
	var (
		name  string
		life  int
		atk   int
		def   int
		magic int
	)

	if lvl <= 5 && lvl > 0 {
		name = "Goblin Soldier"
		life = (lvl * 10)
		atk = 10
		def = 5
		magic = 5
	} else if lvl <= 10 {
		name = "Hob Goblin"
		life = (lvl * 10)
		atk = 15
		def = 10
		magic = 20
	} else if lvl <= 15 {
		name = "Goblin General"
		life = (lvl * 10)
		atk = 20
		def = 15
		magic = 25
	} else if lvl >= 16 {
		name = "Goblin King"
		life = (lvl * 10)
		atk = 25
		def = 20
		magic = 100
	} else {
		name = "Lost Goblin"
		life = 50
		atk = 10
		def = 5
		magic = 5
	}

	return &Monster{Name: name, Life: life, Atk: atk, Def: def, Magic: magic}
}

func NewSlime(lvl int) *Monster {

	var (
		name  string
		life  int
		atk   int
		def   int
		magic int
	)

	if lvl <= 5 && lvl > 0 {
		name = "Blue Slime"
		life = (lvl * 15)
		atk = 10
		def = 5
		magic = 5
	} else if lvl <= 10 {
		name = "Large Slime"
		life = (lvl * 15)
		atk = 15
		def = 5
		magic = 20
	} else if lvl <= 15 {
		name = "Mega Slime"
		life = (lvl * 15)
		atk = 20
		def = 10
		magic = 25
	} else if lvl >= 16 {
		name = "Ultra Slime"
		life = (lvl * 20)
		atk = 25
		def = 15
		magic = 100
	} else {
		name = "Small Slime"
		life = 20
		atk = 5
		def = 5
		magic = 5
	}

	return &Monster{Name: name, Life: life, Atk: atk, Def: def, Magic: magic}
}
