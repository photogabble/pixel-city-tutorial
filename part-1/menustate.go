package main

import "golang.org/x/image/colornames"

type MenuState struct {
	BaseState
}

func (s MenuState) Draw (dt float64) {
	s.g.Window.Clear(colornames.Firebrick)
}

func (s MenuState) Update (dt float64) {
	// ...
}

func (s MenuState) HandleInput() {
	// ...
}

func NewMenuState(g *Game) *MenuState {
	s := MenuState{
		BaseState: BaseState{g},
	}

	return &s
}