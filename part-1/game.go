package main

import (
	"fmt"
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
	fmt.Printf("%+v\n", g.States)
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

		g.PeekState().handleInput()
		g.PeekState().update(dt)
		g.PeekState().draw(dt)
		g.Window.Update()
	}
}