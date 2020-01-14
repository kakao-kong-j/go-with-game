package main

import (
	"testing"

	"github.com/jinhokong/go-with-game/scenemanager"
	"github.com/jinhokong/go-with-game/scenes"
)

func Test_SetScene(t *testing.T) {
	scenemanager.SetScene(&scenes.StartScene{})
}
