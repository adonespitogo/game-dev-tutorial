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
