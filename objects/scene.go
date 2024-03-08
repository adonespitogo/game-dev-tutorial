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
	scene.Archer = &Player{Name: "Marc", Life: 100, Atk: 20, Def: 2, Magic: 1}

	scene.Goblin = NewGoblin(17)
	scene.Slime = &Monster{Name: "Mega Slime", Life: 15, Atk: 10, Def: 2, Magic: 3}

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
			//extra turn or reserved for 2nd monster
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

func (scene *Scene) PrintStatus() {
	fmt.Println(" ")
	fmt.Println("Players")
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Warrior.Name, scene.Warrior.Life, scene.Warrior.Atk, scene.Warrior.Def, scene.Warrior.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Mage.Name, scene.Mage.Life, scene.Mage.Atk, scene.Mage.Def, scene.Mage.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Archer.Name, scene.Archer.Life, scene.Archer.Atk, scene.Archer.Def, scene.Archer.Magic)
	fmt.Println(" ")
	fmt.Println("Monster")
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Goblin.Name, scene.Goblin.Life, scene.Goblin.Atk, scene.Goblin.Def, scene.Goblin.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Slime.Name, scene.Slime.Life, scene.Slime.Atk, scene.Slime.Def, scene.Slime.Magic)
	fmt.Println(" ")
}
