package scenemanager

import "github.com/hajimehoshi/ebiten"

// Scene interface
type Scene interface {
	Startup()
	Update(*ebiten.Image) error
}

type scenemanager struct {
	currentScene Scene
}

var manager *scenemanager

func init() {
	manager = &scenemanager{}
}

func Update(screen *ebiten.Image) error {
	if manager.currentScene != nil {
		return manager.currentScene.Update(screen)
	}
	return nil
}

func SetScene(scene Scene) {
	manager.currentScene = scene
	scene.Startup()
}
