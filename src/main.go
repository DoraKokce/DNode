package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "test 0.1v")
	rl.SetTargetFPS(60)

	scene := makeScene(200, 200)

	for !rl.WindowShouldClose() {
		scene.drawBackground()
	}
	rl.CloseWindow()
}
