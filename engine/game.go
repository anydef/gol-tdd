package engine

import "errors"

type Game struct {
	Side int
}

func (g *Game) GetStatusOf(i int) (bool, error) {
	if i < 0 || i >= g.Side {
		return false, errors.New("Index out of visible range")
	}
	return false, nil
}

func NewGame(side_len int) Game {
	return Game{Side: side_len}
}
