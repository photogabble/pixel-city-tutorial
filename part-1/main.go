package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	game := Game{
		States: NewStack(),
	}

	pixelgl.Run(game.GameLoop)
}