package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	game := Game{
		States: NewStack(),
	}

	game.LoadTextures()
	//game.SpriteManager.Debug()
	game.PushState(NewSpriteSheetTestState(&game))

	pixelgl.Run(game.GameLoop)
}