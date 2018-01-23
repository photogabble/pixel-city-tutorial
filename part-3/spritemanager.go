package main

import (
	"image"
	_ "image/png"
	"image/color"
	"os"
	"github.com/faiface/pixel"
	"image/draw"
	"fmt"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

type Sprite struct {
	Bounds		pixel.Rect
}

type SpriteManager struct {
	sprites		map[string]Sprite
	yOffset		int
	spriteSheet image.Image
}

func (sM SpriteManager) decodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{X: x + offsetX, Y: y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}

// https://stackoverflow.com/questions/35964656/golang-how-to-concatenate-append-images-to-one-another
func (sM *SpriteManager) appendToSpriteSheet(img image.Image) {
	// collect pixel data from each image
	pixels1 := sM.decodePixelsFromImage(sM.spriteSheet, 0, 0)
	// the second image has a Y-offset of sM.spriteSheet's max Y (appended at bottom)
	pixels2 := sM.decodePixelsFromImage(img, 0, sM.spriteSheet.Bounds().Max.Y)

	pixelSum := append(pixels1, pixels2...)

	// Set a new size for the new image equal to the max width
	// of bigger image and max height of two images combined

	// Identify bigger width
	w := sM.spriteSheet.Bounds().Dx()
	h := sM.spriteSheet.Bounds().Dy() + img.Bounds().Dy()

	if img.Bounds().Dx() > w {
		w = img.Bounds().Dx()
	}

	newRect := image.Rectangle{
		Min: image.Point{0,0},
		Max: image.Point{w,h},
	}

	finImage := image.NewRGBA(newRect)
	// This is the cool part, all you have to do is loop through
	// each Pixel and set the image's color on the go
	for _, px := range pixelSum {
		finImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}

	draw.Draw(finImage, finImage.Bounds(), finImage, image.Point{0,0}, draw.Src)
	sM.spriteSheet = finImage
	sM.yOffset += img.Bounds().Dy()
}

func (sM *SpriteManager) bootSpriteSheet(img image.Image) {
	sM.yOffset = img.Bounds().Dy()
	sM.spriteSheet = img
}

func (sM *SpriteManager) LoadTexture(name string, relativePath string) (error) {
	file, err := os.Open(relativePath)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	if sM.spriteSheet == nil {
		sM.bootSpriteSheet(img)
	} else {
		sM.appendToSpriteSheet(img)
	}

	fmt.Println("Added [", name ,"] new spritesheet dimensions W:", sM.spriteSheet.Bounds().Dx(), "H:", sM.spriteSheet.Bounds().Dy())

	sM.sprites[name] = Sprite{
		Bounds: pixel.R(float64(img.Bounds().Min.X), float64(img.Bounds().Min.Y + sM.yOffset), float64(img.Bounds().Max.X), float64(img.Bounds().Max.Y + sM.yOffset)),
	}
	return nil
}

func (sM SpriteManager) GetSpriteSheet() pixel.Picture {
	return pixel.PictureDataFromImage(sM.spriteSheet)
}

func (sM *SpriteManager) Debug() {
	for k, v := range sM.sprites {
		fmt.Printf("%s -> %+v\n", k, v)
	}
}

func NewSpriteManager() *SpriteManager {
	tM := SpriteManager{sprites: make(map[string]Sprite)}
	return &tM
}

