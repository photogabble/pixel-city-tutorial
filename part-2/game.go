package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"time"
)

type Game struct {
	States    *stack
	Window    *pixelgl.Window
}

func (g *Game) PushState(state GameState) {
	g.States.Push(state)
}

func (g *Game) PopState() {
	g.States.Pop()
}

func (g *Game) ChangeState(state GameState) {
	if g.States.Len() > 0 {
		g.States.Pop()
	}
	g.PushState(state)
}

func (g *Game) PeekState() GameState {
	if g.States.Len() == 0 {
		return nil
	}
	return g.States.Peek().(GameState)
}

func (g *Game) GameLoop() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go Build a City",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	g.Window = win
	clock := time.Now()

	for !g.Window.Closed(){
		if g.PeekState() == nil {
			break
		}

		dt := time.Since(clock).Seconds()
		clock = time.Now()

		g.PeekState().HandleInput()
		g.PeekState().Update(dt)
		g.PeekState().Draw(dt)
		g.Window.Update()
	}
}