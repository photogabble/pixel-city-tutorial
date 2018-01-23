package main

import (
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel"
)

type SpriteSheetTestState struct {
	BaseState
	sprite *pixel.Sprite
}

func (s SpriteSheetTestState) Draw (dt float64) {
	s.g.Window.Clear(colornames.Whitesmoke)
	s.sprite.Draw(s.g.Window, pixel.IM.Moved(s.g.Window.Bounds().Center()))
}

func (s SpriteSheetTestState) Update (dt float64) {
	// ...
}

func (s SpriteSheetTestState) HandleInput() {
	// ...
}

func NewSpriteSheetTestState(g *Game) *SpriteSheetTestState {
	spriteSheet := g.SpriteManager.GetSpriteSheet()

	s := SpriteSheetTestState{
		BaseState: BaseState{g},
		sprite: pixel.NewSprite(spriteSheet, spriteSheet.Bounds()),
	}

	return &s
}