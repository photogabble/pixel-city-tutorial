package main

type GameState interface {
	draw(dt float64)
	update(dt float64)
	handleInput()
	setGame(g *Game)
}

type BaseState struct {
	g *Game
}

func (s *BaseState) setGame(g *Game) {
	s.g = g
}