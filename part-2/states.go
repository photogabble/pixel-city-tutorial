package main

type GameState interface {
	Draw(dt float64)
	Update(dt float64)
	HandleInput()
	SetGame(g *Game)
}

type BaseState struct {
	g *Game
}

func (s *BaseState) SetGame(g *Game) {
	s.g = g
}