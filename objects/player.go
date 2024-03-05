package objects

type Player struct {
	Name string
    Life int
	Atk  int
	Def  int
}

// Add a method to the player
func (self *Player) AttackPlayer(target *Player) {
    target.Life = target.Life - self.Atk
}

func (self *Player) AttackMonster(mons *Monster) {
    mons.Life = mons.Life - self.Atk
}

func (self *Player) Heal(pts int) {
    self.Life = self.Life + pts
}

func (self *Player) IsDead() bool {
    return self.Life <= 0
}
