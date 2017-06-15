package main

import (
	"image"
	"image/color"
	"github.com/faiface/pixel"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

type SpriteManager struct {
	sprites		map[string]pixel.Picture
	spriteSheet image.Image
}



func NewSpriteManager() *SpriteManager {
	tM := SpriteManager{sprites: make(map[string]pixel.Picture)}
	return &tM
}

