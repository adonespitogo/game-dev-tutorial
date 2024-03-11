package objects

import (
	"fmt"
	"time"
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

	scene.Mage = NewPlayer("1")
	scene.Warrior = NewPlayer("2")
	scene.Archer = NewPlayer("3")
	scene.IntroStory()
	scene.Goblin = NewGoblin(12)
	scene.Slime = NewSlime(10)
}

func (scene *Scene) Battle(monster *Monster) {
	fmt.Printf("the journey continues\n")
	DotLoad(3)
	Delay(3)
	fmt.Printf("The heroes encountered a %s!\n", monster.Name)
	DotLoad(3)
	DialogSpacer(1)
	fmt.Println("Battle Start")
	DialogSpacer(1)

	for x := 0; !(scene.AllPlayerDead() || monster.IsDead()); x++ {

		if x == 0 { // 1st Turn

			// fmt.Print("1: ")
			scene.Warrior.PlayerTurn(monster)

		} else if x == 1 { // 2nd Turn

			// fmt.Print("2: ")
			scene.Mage.PlayerTurn(monster)

		} else if x == 2 { // 3rd Turn

			// fmt.Print("3: ")
			scene.Archer.PlayerTurn(monster)

		} else if x == 3 { // 4th turn

			// fmt.Print("4: ")
			monster.AttackPlayer(scene.Mage)
			monster.AttackPlayer(scene.Warrior)
			monster.AttackPlayer(scene.Archer)
			fmt.Printf("%s launched an attack\n", monster.Name)
			DialogSpacer(1)

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
	DialogSpacer(1)
	fmt.Printf("Game over!")
}

func (scene *Scene) PrintStatus() {
	DialogSpacer(1)
	fmt.Println("Players")
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Warrior.Name, scene.Warrior.Life, scene.Warrior.Atk, scene.Warrior.Def, scene.Warrior.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Mage.Name, scene.Mage.Life, scene.Mage.Atk, scene.Mage.Def, scene.Mage.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Archer.Name, scene.Archer.Life, scene.Archer.Atk, scene.Archer.Def, scene.Archer.Magic)
	DialogSpacer(1)
	fmt.Println("Monster")
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Goblin.Name, scene.Goblin.Life, scene.Goblin.Atk, scene.Goblin.Def, scene.Goblin.Magic)
	fmt.Printf("%s: Life: %d, Atk: %d, Def: %d, Magic: %d,\n", scene.Slime.Name, scene.Slime.Life, scene.Slime.Atk, scene.Slime.Def, scene.Slime.Magic)
	DialogSpacer(1)
}

func DialogSpacer(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("")
	}
}

func DotLoad(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(600 * time.Millisecond)
	DialogSpacer(1)
}

func Delay(n time.Duration) {
	time.Sleep(n * time.Second)
}

func (s *Scene) IntroStory() {
	DialogSpacer(1)
	fmt.Printf("Welcome %s, %s and %s\n", s.Mage.Name, s.Warrior.Name, s.Archer.Name)
	Delay(2)
	fmt.Printf("Three new hereos now join to embark on a new journey.\n")
	Delay(2)
	fmt.Printf("In a mystical forest,")
	Delay(2)
	fmt.Printf(" three heroes stand ready.\n")
	Delay(2)
	fmt.Printf("Bound by fate, ")
	Delay(2)
	fmt.Printf("they venture forth into the unknown,\n")
	Delay(2)
	fmt.Printf("where challenges await ")
	Delay(2)
	fmt.Printf("and legends are made.")
	DotLoad(7)
	Delay(5)
	DialogSpacer(1)
	DotLoad(3)
}
