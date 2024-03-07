package main

import (
	"fmt"

	"github.com/adonespitogo/game-dev-tutorial/objects"
)

func main() {
	scene := objects.NewScene()
	scene.Start()

	scene.PrintStatus()

	scene.Battle(scene.Slime)

	scene.PrintStatus()

	if scene.AllPlayerDead() {
		fmt.Println("GAME OVER, YOU LOSE")
	} else {
		scene.Battle(scene.Goblin)
	}
	scene.PrintStatus()
}

/*

package main

import "fmt"

type Monster struct {
    Level int
    Attack int
	Life int
}

type Player struct {
	Life int
}

func NewMonster(level int) *Monster {
	return &Monster{Life: 100, Level: level, Attack: level * 100}
}

func NewPlayer() *Player {
	return &Player{Life: 100}
}

func main() {

	// scene 1
	// player := NewPlayer()
	goblin1 := NewMonster(1)
	goblin2 := NewMonster(2)

	// scene 2
    monster = NewMonster()
}
*/
