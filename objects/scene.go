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
