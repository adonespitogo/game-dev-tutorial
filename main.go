package main

import "github.com/adonespitogo/game-dev-tutorial/objects"

func main() {
	scene := objects.NewScene()
	scene.Start()
	scene.Battle(scene.Goblin)
	scene.Battle(scene.Slime)
}
