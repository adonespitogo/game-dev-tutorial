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
	if lvl <= 5 {
		return &Monster{Name: "Goblin Soldier", Life: (lvl * 10), Atk: 10, Def: 5, Magic: 5}
	} else if lvl <= 10 {
		return &Monster{Name: "Hob Goblin", Life: (lvl * 10), Atk: 15, Def: 10, Magic: 20}
	} else if lvl <= 15 {
		return &Monster{Name: "Goblin General", Life: (lvl * 10), Atk: 20, Def: 15, Magic: 50}
	} else if lvl <= 20 {
		return &Monster{Name: "Goblin King", Life: (lvl * 10), Atk: 25, Def: 20, Magic: 100}
	} else {
		return &Monster{Name: "Lost Goblin", Life: 50, Atk: 10, Def: 5, Magic: 5}
	}
}

func NewSlime(lvl int) *Monster {
	if lvl <= 5 {
		return &Monster{Name: "Blue Slime", Life: (lvl * 10), Atk: 10, Def: 5, Magic: 5}
	} else if lvl <= 10 {
		return &Monster{Name: "Large Slime", Life: (lvl * 10), Atk: 15, Def: 5, Magic: 20}
	} else if lvl <= 15 {
		return &Monster{Name: "Mega Slime", Life: (lvl * 15), Atk: 20, Def: 10, Magic: 50}
	} else if lvl <= 20 {
		return &Monster{Name: "Ultra Slime", Life: (lvl * 20), Atk: 25, Def: 15, Magic: 100}
	} else {
		return &Monster{Name: "Small Slime", Life: 20, Atk: 10, Def: 5, Magic: 5}
	}
}
