package objects

import "fmt"

func NewScene() *Scene {
	return &Scene{}
}

type Scene struct {
	Mage    *Player
	Warrior *Player
	Monster *Monster
}

func (scene *Scene) Start() {
	scene.Mage = &Player{Name: "Mage", Life: 10, Atk: 5}

	monster1 := &Monster{Name: "Monster 1", Life: 10, Atk: 5, Def: 2}
	player1 := &Player{Name: "Player 1", Life: 1, Atk: 10}

	fmt.Printf("Monster life: %d\n", monster1.Life)

	player1.AttackMonster(monster1)

	if player1.IsDead() {
		fmt.Println("Player 1 is dead")
	}

	fmt.Printf("Monster life after attack: %d\n", monster1.Life)

	scene.End()
}

func (scene *Scene) End() {
	fmt.Printf("Game over! Mage Life: %d", scene.Mage.Life)
}

func Battle(p1 *Player, p2 *Player, p3 *Monster) {

	x := 0
	y := 0

	fmt.Println("Battle Start")

	for !(p1.IsDead() && p2.IsDead() || p3.IsDead()) {

		if x > y { // 2nd Turn
			fmt.Print("2: ")

			if !p2.IsDead() {
				p2.AttackMonster(p3)
				fmt.Printf("%s attacked %s", p2.Name, p3.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			y++

		} else if x%2 == 0 { // 1st Turn
			fmt.Print("1: ")

			if !p1.IsDead() {
				p1.AttackMonster(p3)
				fmt.Printf("%s attacked %s", p1.Name, p3.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			x++

		} else { // 3rd turn
			fmt.Print("3: ")

			p3.AttackPlayer(p1)
			p3.AttackPlayer(p2)
			fmt.Printf("%s launched an attack", p3.Name)

			x++
			y++

		}

	}

	fmt.Println("Battle End")

}
