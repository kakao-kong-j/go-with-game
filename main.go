package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/jinhokong/go-with-game/global"
	"github.com/jinhokong/go-with-game/scenemanager"
	"github.com/jinhokong/go-with-game/scenes"
)

func main() {
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
