package objects

import (
	"fmt"
)

func NewScene() *Scene {
	return &Scene{}
}

type Scene struct {
	Mage    *Player
	Warrior *Player
	Goblin  *Monster
	Slime   *Monster
	Archer  *Player
}

func (scene *Scene) Start() {

	scene.Mage = &Player{Name: "Kyle", Life: 100, Atk: 10, Def: 5, Magic: 25}
	scene.Warrior = &Player{Name: "Carlo", Life: 150, Atk: 20, Def: 10, Magic: 10}
	scene.Goblin = &Monster{Name: "Goblin King", Life: 50, Atk: 1500, Def: 500, Magic: 5}
	scene.Slime = &Monster{Name: "Mega Slime", Life: 15, Atk: 10, Def: 2, Magic: 3}
	scene.Archer = &Player{Name: "Marc", Life: 100, Atk: 20, Def: 2, Magic: 1}

}

func (scene *Scene) Battle(monster *Monster) {

	fmt.Println("Battle Start")

	for x := 0; !(scene.AllPlayerDead() || monster.IsDead()); x++ {

		if x == 0 { // 1st Turn

			fmt.Print("1: ")
			scene.Warrior.PlayerTurn(monster)

		} else if x == 1 { // 2nd Turn

			fmt.Print("2: ")
			scene.Mage.PlayerTurn(monster)

		} else if x == 2 { // 3rd Turn

			fmt.Print("3: ")
			scene.Archer.PlayerTurn(monster)

		} else if x == 3 { // 4th turn

			fmt.Print("4: ")
			monster.AttackPlayer(scene.Mage)
			monster.AttackPlayer(scene.Warrior)
			monster.AttackPlayer(scene.Archer)
			fmt.Printf("%s launched an attack\n", monster.Name)

		} else {
			//extra turn or reserved
			x = -1
		}
	}
	//Dodongcaloy was here
	fmt.Println("Battle End")
}

func (scene *Scene) AllPlayerDead() bool {
	return scene.Mage.IsDead() && scene.Warrior.IsDead() && scene.Archer.IsDead()
}

func (scene *Scene) AllMonsterDead() bool {
	return scene.Goblin.IsDead() && scene.Slime.IsDead()
}

func (scene *Scene) End() {
	fmt.Printf("Game over! Mage Life: %d\n", scene.Mage.Life)
}
