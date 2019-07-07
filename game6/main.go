package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tuckersGo/12janggi/game6/global"
	"github.com/tuckersGo/12janggi/game6/scenemanager"
	"github.com/tuckersGo/12janggi/game6/scenes"
)

func main() {
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
