package main

import (
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
)

type SpriteSheetTestState struct {
	BaseState
	sprites []pixel.Sprite
	sP      int
}

func (s SpriteSheetTestState) Draw (dt float64) {
	s.g.Window.Clear(colornames.Whitesmoke)
	s.sprites[s.sP].Draw(s.g.Window, pixel.IM.Moved(s.g.Window.Bounds().Center()))
}

func (s SpriteSheetTestState) Update (dt float64) {
	// ...
}

func (s *SpriteSheetTestState) HandleInput() {
	if s.g.Window.JustPressed(pixelgl.KeyEnter) {
		if (s.sP + 1) > len(s.sprites) - 1 {
			s.sP = 0
		} else {
			s.sP++
		}

		t := s.sprites[s.sP]
		fmt.Println("Sprite Pointer [", s.sP,"] (", t.Frame().Min.X, ",",t.Frame().Min.Y,"), (",t.Frame().Max.X,",",t.Frame().Max.Y,")")
	}
}

func NewSpriteSheetTestState(g *Game) *SpriteSheetTestState {
	spriteSheet := g.SpriteManager.GetSpriteSheet()

	var tileTypes = []pixel.Sprite {
		*pixel.NewSprite(spriteSheet, spriteSheet.Bounds()),
		g.SpriteManager.GetSprite("grass"),
		g.SpriteManager.GetSprite("forest"),
		g.SpriteManager.GetSprite("water"),
		g.SpriteManager.GetSprite("residential"),
		g.SpriteManager.GetSprite("commercial"),
		g.SpriteManager.GetSprite("industrial"),
		g.SpriteManager.GetSprite("road"),
	}

	s := SpriteSheetTestState{
		BaseState: BaseState{g},
		sprites:   tileTypes,
		sP:        0,
	}

	return &s
}