package main

import (
	"fmt"

	"github.com/adonespitogo/game-dev-tutorial/objects"
)

func main() {

	monsters := []*objects.Monster{
		objects.NewGoblin(1),
		objects.NewGoblin(2),
		objects.NewGoblin(3),
	}

	mage := objects.NewPlayer("1")
	warrior := objects.NewPlayer("2")
	archer := objects.NewPlayer("3")

	for _, monster := range monsters {
		scene := objects.NewScene(mage, warrior, archer, monster)
		scene.Start()
		if scene.AllPlayerDead() {
			fmt.Println("GAME OVER, YOU LOSE")
		}
		scene.PrintStatus()
	}
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
