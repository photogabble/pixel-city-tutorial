package main

type RedState struct {
	BaseState
}

func (s RedState) draw (dt float64) {
	// ...
}

func (s RedState) update (dt float64) {
	// ...
}

func (s RedState) handleInput() {
	// ...
}

func NewRedState(g *Game) *RedState {
	s := RedState{
		BaseState: BaseState{g},
	}

	return &s
}