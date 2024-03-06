package objects

import "fmt"

func NewScene() *Scene {
	return &Scene{}
}

type Scene struct {
	Mage    *Player
	Warrior *Player
	Goblin  *Monster
	Slime   *Monster
}

func (scene *Scene) Start() {

	scene.Mage = &Player{Name: "Kyle", Life: 100, Atk: 10, Def: 5, Magic: 25}
	scene.Warrior = &Player{Name: "Carlo", Life: 150, Atk: 20, Def: 10, Magic: 10}
	scene.Goblin = &Monster{Name: "Gobs", Life: 50, Atk: 15, Def: 5, Magic: 5}
	scene.Slime = &Monster{Name: "Gobs", Life: 15, Atk: 10, Def: 2, Magic: 3}

	fmt.Printf("Monster life: %d\n", scene.Goblin.Life)

	scene.Warrior.AttackMonster(scene.Goblin)

	if scene.Mage.IsDead() {
		fmt.Println("Mage is dead")
	}

	fmt.Printf("Goblin life after attack: %d\n", scene.Goblin.Life)

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

		if x%2 == 0 { // 1st Turn
			fmt.Print("2: ")

			if !scene.Warrior.IsDead() {
				scene.Warrior.AttackMonster(monster)
				fmt.Printf("%s attacked %s", scene.Warrior.Name, monster.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			x++

		} else if x > y { // 2nd Turn
			fmt.Print("1: ")

			if !scene.Mage.IsDead() {
				scene.Mage.AttackMonster(monster)
				fmt.Printf("%s attacked %s", scene.Mage.Name, monster.Name)
			} else {
				fmt.Println("Player Dead - Skip Turn")
			}

			y++

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
