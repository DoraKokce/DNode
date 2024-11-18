package dnode

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene struct {
	SceneWidth  int32
	SceneHeight int32
}

func (scn Scene) drawBackground() {
	h := scn.SceneHeight / 2
	w := scn.SceneWidth / 2
	rl.DrawRectangle(-w, -h, w, h, rl.NewColor(30, 30, 30, 255))
}

func makeScene(width int32, height int32) Scene {
	return Scene{
		SceneWidth:  width,
		SceneHeight: height,
	}
}
