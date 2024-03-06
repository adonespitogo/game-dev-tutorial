package objects

import "fmt"

func NewScene() *Scene {
	return &Scene{}
}

type Scene struct {
	Mage    *Player
	Warrior *Player
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

func (scene *Scene) Battle(monster *Monster) {

	x := 0
	y := 0

	fmt.Println("Battle Start")

	for !(scene.Mage.IsDead() && scene.Warrior.IsDead() || monster.IsDead()) {

		if x > y { // 2nd Turn
			fmt.Print("2: ")

			if !scene.Warrior.IsDead() {
				scene.Warrior.AttackMonster(monster)
				fmt.Printf("%s attacked %s", scene.Warrior.Name, monster.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			y++

		} else if x%2 == 0 { // 1st Turn
			fmt.Print("1: ")

			if !scene.Mage.IsDead() {
				scene.Mage.AttackMonster(monster)
				fmt.Printf("%s attacked %s", scene.Mage.Name, monster.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			x++

		} else { // 3rd turn
			fmt.Print("3: ")

			monster.AttackPlayer(scene.Mage)
			monster.AttackPlayer(scene.Warrior)
			fmt.Printf("%s launched an attack", monster.Name)

			x++
			y++

		}

	}

	fmt.Println("Battle End")

}
