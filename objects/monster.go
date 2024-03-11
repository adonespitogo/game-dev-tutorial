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
		return &Monster{Name: "Goblin General", Life: (lvl * 10), Atk: 20, Def: 15}
	} else if lvl >= 16 {
		return &Monster{Name: "Goblin King", Life: (lvl * 10), Atk: 25, Def: 20, Magic: 100}
	} else {
		return &Monster{Name: "Lost Goblin", Life: 50, Atk: 10, Def: 5, Magic: 5}
	}

	return &Monster{Name: name, Life: life, Atk: atk, Def: def, Magic: magic}
}

func NewSlime(lvl int) *Monster {
	if lvl <= 5 && lvl > 0 {
		return &Monster{Name: "Blue Slime", Life: (lvl * 10), Atk: 10, Def: 5, Magic: 5}
	} else if lvl <= 10 {
		return &Monster{Name: "Large Slime", Life: (lvl * 10), Atk: 15, Def: 5, Magic: 20}
	} else if lvl <= 15 {
		return &Monster{Name: "Mega Slime", Life: (lvl * 15), Atk: 20, Def: 10, Magic: 50}
	} else if lvl >= 16 {
		return &Monster{Name: "Ultra Slime", Life: (lvl * 20), Atk: 25, Def: 15, Magic: 100}
	} else {
		return &Monster{Name: "Small Slime", Life: 20, Atk: 10, Def: 5, Magic: 5}
	}
}
